// +build cgo

package codec2

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/go-voice/voice/internal/tests"
)

func TestMode(t *testing.T) {
	for mode := Mode3200; mode <= Mode700C; mode++ {
		t.Run(mode.String(), func(t *testing.T) {
			codec, err := New(mode)
			if err != nil {
				t.Fatal(err)
			}
			defer codec.Close()

			if l := codec.DecodeBlockSize(); l == 0 {
				t.Fatal("decode block size returned 0")
			} else {
				t.Logf("decode block size %d", l)
			}
			if l := codec.EncodeBlockSize(); l == 0 {
				t.Fatal("encode block size returned 0")
			} else {
				t.Logf("encode block size %d", l)
			}
		})
	}
}

var (
	testModes  = []Mode{Mode3200, Mode1600, Mode1200, Mode700C}
	testSuites = []string{"female", "male", "violin"}
)

func TestDecode(t *testing.T) {
	for _, suite := range testSuites {
		t.Run(suite, func(t *testing.T) {
			for _, mode := range testModes {
				t.Run(mode.String(), func(t *testing.T) {
					testDecode(t, suite, mode)
				})
			}
		})
	}
}

func testDecode(t *testing.T, suite string, mode Mode) {
	t.Helper()

	codec, err := New(mode)
	if err != nil {
		t.Fatal(err)
	}
	defer codec.Close()

	var (
		bytes = tests.LoadBytes(t, filepath.Join("testdata", fmt.Sprintf("%s-%s.codec2", suite, mode)))
		want  = tests.LoadPCMS16(t, filepath.Join("testdata", fmt.Sprintf("%s-%s.decoded.raw", suite, mode)))
		test  = make([]int16, len(want))
	)
	if err = codec.Decode(test, bytes); err != nil {
		t.Fatal(err)
	}
	// tests.ComparePCMS16(t, test, want, 0)
}

func TestEncode(t *testing.T) {
	for _, suite := range testSuites {
		t.Run(suite, func(t *testing.T) {
			for _, mode := range testModes {
				t.Run(mode.String(), func(t *testing.T) {
					testEncode(t, suite, mode)
				})
			}
		})
	}
}

func testEncode(t *testing.T, suite string, mode Mode) {
	t.Helper()

	codec, err := New(mode)
	if err != nil {
		t.Fatal(err)
	}
	defer codec.Close()

	var (
		samples = tests.LoadPCMS16(t, filepath.Join("..", "testdata", fmt.Sprintf("%s.raw", suite)))
		want    = tests.LoadBytes(t, filepath.Join("testdata", fmt.Sprintf("%s-%s.codec2", suite, mode)))
		test    = make([]byte, len(want))
	)
	if blockSize := mode.SamplesPerFrame(); len(samples)%blockSize != 0 {
		samples = samples[:len(samples)-(len(samples)%blockSize)]
	}
	if err = codec.Encode(test, samples); err != nil {
		t.Fatal(err)
	}
	tests.CompareBytes(t, test, want, 0)
}
