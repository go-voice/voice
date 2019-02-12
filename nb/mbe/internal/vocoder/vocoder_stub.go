// +build !cgo

package vocoder

import (
	"errors"

	"github.com/go-voice/voice"
)

func isAvailable() bool {
	return false
}

var errUnavailable = errors.New("mbe: voice is not built with cgo support; vocoder unavailable")

// Vocoder placeholder.
type Vocoder struct {
}

func New(tag voice.Tag, _ int) *Vocoder {
	return &Vocoder{}
}

func (Vocoder) Close() error                     { return errUnavailable }
func (Vocoder) Decode(_ []int16, _ []byte) error { return errUnavailable }
func (Vocoder) DecodeBlockSize() int             { return 0 }
func (Vocoder) DecodeLen(_ int) int              { return 0 }
func (Vocoder) Encode(_ []byte, _ []int16) error { return errUnavailable }
func (Vocoder) EncodeLen(_ int) int              { return 0 }
func (Vocoder) EncodeBlockSize() int             { return 0 }
func (Vocoder) Format() voice.Format             { return voice.Format{Channels: 1, Rate: 8000} }
func (Vocoder) Reset()                           {}
