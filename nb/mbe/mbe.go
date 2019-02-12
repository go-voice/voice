/*
Package mbe implements Multi-Band Excitation codecs (AMBE and IMBE).

Patent notice

This source code is provided for educational purposes only.  It is
a written description of how certain voice encoding/decoding
algorythims could be implemented.  Executable objects compiled or
derived from this package may be covered by one or more patents.
Readers are strongly advised to check for any patent restrictions or
licencing requirements before compiling or using this source code.


Copyright

The AMBE and IMBE decoders are taken from the mbelib project and are
© Copyright 2010 mbelib Author (MIT license).

The AMBE and IMBE encoders are taken from the op25 project and are
© Copyright 2016 Max H. Parke KA1RBI (GPLv3+ license).
*/
package mbe

import (
	"github.com/go-voice/voice"
	"github.com/go-voice/voice/nb"
	"github.com/go-voice/voice/nb/mbe/internal/vocoder"
)

// DefaultQuality is the default quality for our codecs.
const DefaultQuality = 3

// IsAvailable checks if dv was built with mbe support.
func IsAvailable() bool {
	return vocoder.IsAvailable()
}

func init() {
	/*
		dv.Register(dv.AMBE3600x2400, func() nb.Codec {
			return NewAMBE3600x2400(DefaultQuality)
		})
		dv.Register(dv.AMBE3600x2450, func() nb.Codec {
			return NewAMBE3600x2450(DefaultQuality)
		})
		dv.Register(dv.IMBE7100x4400, func() nb.Codec {
			return NewIMBE7100x4400(DefaultQuality)
		})
		dv.Register(dv.IMBE7200x4400, func() nb.Codec {
			return NewIMBE7200x4400(DefaultQuality)
		})
	*/
}

// MBE is the basis for all Multi-Band Excitation codecs.
type mbe struct {
	coder *vocoder.Vocoder
	tag   voice.Tag
}

// NewAMBE3600x2400 returns an AMBE2+ codec at 2400bps.
func NewAMBE3600x2400(quality int) nb.CodecCloser {
	return newMBE(quality, voice.AMBE3600x2400)
}

// NewAMBE3600x2450 returns an AMBE2+ codec at 2450bps.
func NewAMBE3600x2450(quality int) nb.CodecCloser {
	return newMBE(quality, voice.AMBE3600x2450)
}

// NewIMBE7100x4400 returns an IMBE codec with 7100bps voice.
func NewIMBE7100x4400(quality int) nb.CodecCloser {
	return newMBE(quality, voice.IMBE7100x4400)
}

// NewIMBE7200x4400 returns an IMBE codec with 7200bps voice.
func NewIMBE7200x4400(quality int) nb.CodecCloser {
	return newMBE(quality, voice.IMBE7200x4400)
}

// newMBE returns a new MBE codec with the specified quality.
func newMBE(quality int, tag voice.Tag) *mbe {
	return &mbe{
		coder: vocoder.New(tag, quality),
		tag:   tag,
	}
}

// Format specifier.
func (mbe *mbe) Format() voice.Format {
	return voice.Format{
		Channels: 1,
		Rate:     8000,
	}
}

// Close releases allocated memory.
func (mbe *mbe) Close() error {
	if mbe.coder != nil {
		mbe.coder.Close()
		mbe.coder = nil
	}
	return nil
}

// Decode a buffer
func (mbe *mbe) Decode(dst []int16, src []byte) error {
	return mbe.coder.Decode(dst, src)
}

// Encode a buffer
func (mbe *mbe) Encode(dst []byte, src []int16) error {
	return mbe.coder.Encode(dst, src)
}

// DecodeBlockSize is 160 samples per block.
func (mbe *mbe) DecodeBlockSize() int {
	return 160
}

// DecodedLen returns how many samples can be extracted from n bytes of input.
func (mbe *mbe) DecodedLen(n int) int {
	blockSize := mbe.EncodeBlockSize()
	if n%blockSize != 0 {
		return 0
	}
	return (n / blockSize) * mbe.DecodeBlockSize()
}

// EncodeBlockSize is 8 AMBE bytes or 12 IMBE bytes per block.
func (mbe *mbe) EncodeBlockSize() int {
	return mbe.coder.EncodeBlockSize()
}

// EncodedLen returns how many bytes are required to encode n samples of
// input. If n is not an accepted multiple of block size, 0 is returned.
func (mbe *mbe) EncodedLen(n int) int {
	blockSize := mbe.DecodeBlockSize()
	if n%blockSize != 0 {
		return 0
	}
	return (n / blockSize) * mbe.EncodeBlockSize()
}

// Reset coder state.
func (mbe *mbe) Reset() {
	if mbe.coder != nil {
		mbe.coder.Reset()
	}
}
