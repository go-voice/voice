package voice

import "time"

// Codec is a coder-decoder of voice data.
type Codec interface {
	// DecodeBlockSize is the recommended block size for this decoder. If the
	// block size is not critical, this method will return 0.
	DecodeBlockSize() int

	// DecodedLen returns how many samples can be extracted from n bytes of input.
	DecodedLen(n int) int

	// EncodeBlockSize is the recommended block size for this encoder. If the
	// block size is not critical, this method will return 0.
	EncodeBlockSize() int

	// EncodedLen returns how many bytes are required to encode n samples of
	// input. If n is not an accepted multiple of block size, 0 is returned.
	EncodedLen(n int) int

	// Format of this codec.
	Format() Format

	// Reset the coder state.
	Reset()
}

// Format describes the sample format.
type Format struct {
	// Rate is the number of samples per second.
	Rate Rate

	// Channels is the number of channels. The value of 1 is mono, the value of 2 is stereo.
	// The samples should always be interleaved.
	Channels int

	// Precision is the number of bytes used to encode a single sample.
	Precision int
}

// Rate in Hertz.
type Rate int

// Duration of n samples.
func (rate Rate) Duration(n int) time.Duration {
	return (time.Second * time.Duration(n)) / time.Duration(rate)
}

// Samples is the number of samples during d.
func (rate Rate) Samples(d time.Duration) int {
	return int((d * time.Duration(rate)) / time.Second)
}
