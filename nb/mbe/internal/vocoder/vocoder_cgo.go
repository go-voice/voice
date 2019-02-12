// +build cgo

package vocoder

/*
#cgo CFLAGS: -I. -Iinclude
#cgo LDFLAGS: -lm
#include "vocoder.h"
*/
import (
	"C"
)

import (
	"fmt"
	"unsafe"

	"github.com/go-voice/voice"
)

func isAvailable() bool {
	return true
}

// Vocoder for Multi-Band Excited codecs.
type Vocoder struct {
	ptr C.vocoder
	tag voice.Tag
}

// New vocoder for the desired format.
func New(tag voice.Tag, quality int) *Vocoder {
	ptr := C.vocoder_new(C.int(quality))
	switch tag {
	// DMR ambe
	case voice.AMBE3600x2400:
		C.vocoder_ambe_mode_dstar(ptr)
	case voice.AMBE3600x2450:
		C.vocoder_ambe_mode_dmr(ptr)
	case voice.IMBE7100x4400:
		// TODO
	case voice.IMBE7200x4400:
		// TODO
	}
	return &Vocoder{
		ptr: ptr,
		tag: tag,
	}
}

func (coder *Vocoder) Close() error {
	if coder.ptr != nil {
		C.vocoder_destroy(coder.ptr)
		coder.ptr = nil
	}
	return nil
}

func (coder *Vocoder) Decode(dst []int16, src []byte) error {
	var (
		dstLen       = len(dst)
		dstBlockSize = 160
		srcLen       = len(src)
		srcBlockSize = coder.EncodeBlockSize()
		minLen       = (srcLen / srcBlockSize) * dstBlockSize
	)
	if srcLen%srcBlockSize != 0 {
		return fmt.Errorf("mbe: src length of %d must be multiple of %d", srcLen, srcBlockSize)
	}
	if dstLen < minLen {
		return fmt.Errorf("mbe: dst length of %d is too small, expected at least %d", dstLen, minLen)
	}

	switch coder.tag {
	case voice.AMBE3600x2400, voice.AMBE3600x2450:
		for srcOffset, dstOffset := 0, 0; srcOffset < srcLen; srcOffset, dstOffset = srcOffset+srcBlockSize, dstOffset+dstBlockSize {
			C.vocoder_ambe_decode(coder.ptr,
				(*C.int16_t)(unsafe.Pointer(&dst[dstOffset])),
				(*C.uint8_t)(unsafe.Pointer(&src[srcOffset])))
		}
	}

	return nil
}

func (coder *Vocoder) Encode(dst []byte, src []int16) error {
	var (
		dstLen       = len(dst)
		dstBlockSize = coder.EncodeBlockSize()
		srcLen       = len(src)
		srcBlockSize = 160
		minLen       = (srcLen / srcBlockSize) * dstBlockSize
	)
	if srcLen%srcBlockSize != 0 {
		return fmt.Errorf("mbe: src length of %d must be multiple of %d", srcLen, srcBlockSize)
	}
	if dstLen < minLen {
		return fmt.Errorf("mbe: dst length of %d is too small, expected at least %d", dstLen, minLen)
	}

	switch coder.tag {
	case voice.AMBE3600x2400, voice.AMBE3600x2450:
		for srcOffset, dstOffset := 0, 0; srcOffset < srcLen; srcOffset, dstOffset = srcOffset+srcBlockSize, dstOffset+dstBlockSize {
			C.vocoder_ambe_encode(coder.ptr,
				(*C.uint8_t)(unsafe.Pointer(&dst[dstOffset])),
				(*C.int16_t)(unsafe.Pointer(&src[srcOffset])))
		}
	case voice.IMBE7100x4400, voice.IMBE7200x4400:
		/*
			for srcOffset, dstOffset := 0, 0; srcOffset < srcLen; srcOffset, dstOffset = srcOffset+srcBlockSize, dstOffset+dstBlockSize {
				C.vocoder_imbe_encode(coder.ptr,
					(*C.int16_t)(unsafe.Pointer(&dst[dstOffset])),
					(*C.uint8_t)(unsafe.Pointer(&src[srcOffset])))
			}
		*/
	}

	return nil
}

// EncodeBlockSize is the recommended block size for this encoder. If the
// block size is not critical, this method will return 0.
func (coder *Vocoder) EncodeBlockSize() int {
	switch coder.tag {
	case voice.AMBE3600x2400:
		return 9 // 72 bits per frame
	case voice.AMBE3600x2450:
		return 7 // 49 bits per frame
	case voice.IMBE7100x4400:
		return 0 // TODO
	case voice.IMBE7200x4400:
		return 0 // TODO
	default:
		return 0
	}
}

func (coder *Vocoder) Reset() {
	C.vocoder_reset(coder.ptr)
}
