package tests

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"
	"testing"
)

// CompareBytes compares two byte buffers.
func CompareBytes(t *testing.T, test, want []byte, epsilon byte) {
	t.Helper()

	var (
		testLen = len(test)
		wantLen = len(want)
	)
	if testLen != wantLen {
		t.Fatalf("expected %d bytes, got %d", wantLen, testLen)
	}
	for i, value := range want {
		delta := test[i] - value
		if delta < 0 {
			delta = -delta
		}
		if delta > epsilon {
			t.Fatalf("expected byte value %#02x at %d, got %#02x", value, i, test[i])
		}
	}
}

// LoadBytes ...
func LoadBytes(t *testing.T, name string) (buf []byte) {
	t.Helper()

	var err error
	if buf, err = ioutil.ReadFile(name); err != nil {
		if os.IsNotExist(err) {
			t.Skip(err)
		}
		t.Fatal(err)
	}

	return
}

// ComparePCMS16 compares two sample buffers.
func ComparePCMS16(t *testing.T, test, want []int16, epsilon int16) {
	t.Helper()

	var (
		testLen = len(test)
		wantLen = len(want)
	)
	if testLen != wantLen {
		t.Fatalf("expected %d samples, got %d", wantLen, testLen)
	}
	for i, sample := range want {
		delta := test[i] - sample
		if delta < 0 {
			delta = -delta
		}
		if delta > epsilon {
			t.Fatalf("expected sample value %d at %d, got %d", sample, i, test[i])
		}
	}
}

// LoadPCMS16 test samples.
func LoadPCMS16(t *testing.T, name string) (pcm []int16) {
	t.Helper()

	buf := LoadBytes(t, name)
	if len(buf)&1 == 1 {
		t.Fatalf("%s: odd number of bytes", name)
	}

	pcm = make([]int16, len(buf)>>1)
	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &pcm); err != nil {
		t.Fatalf("%s: error reading samples: %v", name, err)
	}

	return
}
