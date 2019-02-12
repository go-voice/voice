// +build cgo

package codec2

/*
#cgo CFLAGS: -I. -Wno-absolute-value
#cgo LDFLAGS: -lm

#include "codec2.h"
typedef struct CODEC2 codec2;

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/go-voice/voice"

	"github.com/go-voice/voice/nb"
)

// Mode setting
type Mode C.int

var (
	modeName = map[Mode]string{
		Mode3200: "3200",
		Mode2400: "2400",
		Mode1600: "1600",
		Mode1400: "1400",
		Mode1300: "1300",
		Mode1200: "1200",
		Mode700:  "700",
		Mode700B: "700B",
		Mode700C: "700C",
	}
	modeBitsPerFrame = map[Mode]int{
		Mode3200: 64,
		Mode2400: 48,
		Mode1600: 64,
		Mode1400: 56,
		Mode1300: 52,
		Mode1200: 48,
		Mode700:  28,
		Mode700B: 28,
		Mode700C: 28,
	}
	modeSamplesPerFrame = map[Mode]int{
		Mode3200: 160,
		Mode2400: 160,
		Mode1600: 320,
		Mode1400: 320,
		Mode1300: 320,
		Mode1200: 320,
		Mode700:  320,
		Mode700B: 320,
		Mode700C: 320,
	}
)

func (mode Mode) BitsPerFrame() int {
	if n, ok := modeBitsPerFrame[mode]; ok {
		return n
	}
	return 0
}

func (mode Mode) BytesPerFrame() int {
	if n, ok := modeBitsPerFrame[mode]; ok {
		return (n + 7) >> 3
	}
	return 0
}

func (mode Mode) SamplesPerFrame() int {
	if n, ok := modeSamplesPerFrame[mode]; ok {
		return n
	}
	return 0
}

func (mode Mode) String() string {
	if s, ok := modeName[mode]; ok {
		return s
	}
	return "invalid"
}

// Supported modes
const (
	Mode3200 Mode = C.CODEC2_MODE_3200
	Mode2400 Mode = C.CODEC2_MODE_2400
	Mode1600 Mode = C.CODEC2_MODE_1600
	Mode1400 Mode = C.CODEC2_MODE_1400
	Mode1300 Mode = C.CODEC2_MODE_1300
	Mode1200 Mode = C.CODEC2_MODE_1200
	Mode700  Mode = C.CODEC2_MODE_700
	Mode700B Mode = C.CODEC2_MODE_700B
	Mode700C Mode = C.CODEC2_MODE_700C
)

type Codec2 struct {
	ptr  *C.codec2
	mode Mode
}

func New(mode Mode) (nb.Codec, error) {
	if mode < Mode3200 || mode > Mode700C {
		return nil, errors.New("codec2: unsupported mode")
	}

	ptr := C.codec2_create(C.int(mode))
	C.codec2_set_natural_or_gray(ptr, 1)

	return &Codec2{
		ptr:  ptr,
		mode: mode,
	}, nil
}

func (codec *Codec2) Close() error {
	if codec.ptr != nil {
		C.codec2_destroy(codec.ptr)
		codec.ptr = nil
	}
	return nil
}

func (codec *Codec2) Decode(dst []int16, src []byte) error {
	var (
		srcLen    = len(src)
		srcStep   = codec.mode.BytesPerFrame()
		srcOffset int
		dstLen    = len(dst)
		dstStep   = codec.mode.SamplesPerFrame()
		dstOffset int
		minLen    = codec.DecodedLen(srcLen)
	)
	if srcLen == 0 && dstLen == 0 {
		return nil
	} else if minLen == 0 {
		return fmt.Errorf("codec2: source buffer size of %d bytes does not align with the block size", srcLen)
	} else if dstLen < minLen {
		return fmt.Errorf("codec2: output buffer size of %d is too small to hold %d decoded samples", dstLen, minLen)
	}

	//log.Printf("decode %d step %d (%d frames) to %d step %d", srcLen, srcStep, srcLen/srcStep, dstLen, dstStep)

	for srcOffset < srcLen {
		C.codec2_decode(codec.ptr,
			(*C.short)(unsafe.Pointer(&dst[dstOffset])),
			(*C.uchar)(unsafe.Pointer(&src[srcOffset])))

		srcOffset += srcStep
		dstOffset += dstStep
	}

	return nil
}

func (codec *Codec2) DecodeBlockSize() int {
	return codec.mode.BytesPerFrame()
}

// DecodedLen returns how many samples can be extracted from n bytes of input.
func (codec *Codec2) DecodedLen(n int) int {
	var (
		samples = codec.mode.SamplesPerFrame()
		bytes   = codec.mode.BytesPerFrame()
	)
	if n == 0 || n%bytes != 0 {
		return 0
	}
	return (n * samples) / bytes
}

func (codec *Codec2) EncodeBlockSize() int {
	return codec.mode.SamplesPerFrame()
}

// EncodedLen returns how many bytes are required to encode n samples of input.
func (codec *Codec2) EncodedLen(n int) int {
	var (
		samples = codec.mode.SamplesPerFrame()
		bytes   = codec.mode.BytesPerFrame()
	)
	if n == 0 || n%samples != 0 {
		return 0
	}
	return (n * bytes) / samples
}

func (codec *Codec2) Encode(dst []byte, src []int16) error {
	var (
		srcLen    = len(src)
		srcStep   = codec.mode.SamplesPerFrame()
		srcOffset int
		dstLen    = len(dst)
		dstStep   = codec.mode.BytesPerFrame()
		dstOffset int
		minLen    = codec.EncodedLen(srcLen)
	)
	if srcLen == 0 && dstLen == 0 {
		return nil
	} else if minLen == 0 {
		return fmt.Errorf("codec2: source buffer size of %d samples does not align with the block size", srcLen)
	} else if dstLen < minLen {
		return fmt.Errorf("codec2: output buffer size of %d is too small to hold %d encoded samples", dstLen, minLen)
	}

	//log.Printf("encode %d step %d (%d frames) to %d step %d", srcLen, srcStep, srcLen/srcStep, dstLen, dstStep)

	for srcOffset < srcLen {
		C.codec2_encode(codec.ptr,
			(*C.uchar)(unsafe.Pointer(&dst[dstOffset])),
			(*C.short)(unsafe.Pointer(&src[srcOffset])))

		dstOffset += dstStep
		srcOffset += srcStep
	}

	return nil
}

func (Codec2) Format() voice.Format {
	return voice.Format{
		Channels: 1,
		Rate:     8000,
	}
}

func (Codec2) Reset() {}

var _ nb.Codec = (*Codec2)(nil)
