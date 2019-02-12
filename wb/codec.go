package wb

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/go-voice/voice"
)

// Codec is a coder-decoder of voice data at 16000 Hz.
type Codec interface {
	voice.Codec

	// Encode src samples into buffer dst. If src is not a multiple of the
	// accepted EncodeBlockSize, an error shall be returned. If dst is too
	// small to encode len(src) samples, an error shall be returned.
	//
	// Codecs shall not retain src nor dst but may use src or dst as scratch
	// space during the call.
	Encode(dst []byte, src []int32) error

	// Decode src bytes into sample slice dst. If src is not a multiple of the
	// accepted DecodeBlockSize, an error shall be returned. If dst is too
	// small to decode all samples contained in src, an error shall be returned.
	//
	// Codecs shall not retain src nor dst but may use src or dst as scratch
	// space during the call.
	Decode(dst []int32, src []byte) error
}

// checkEncodeBounds does boundary checks for the given buffer sizes.
func checkEncodeBounds(codec Codec, dst []byte, src []int32) error {
	var (
		srcLen = len(src)
		dstLen = len(dst)
		minLen = codec.EncodedLen(srcLen)
	)
	if srcLen == 0 && dstLen == 0 {
		return nil
	} else if minLen == 0 {
		return fmt.Errorf("voice: source buffer size of %d samples does not align with the block size", srcLen)
	} else if dstLen < minLen {
		return fmt.Errorf("voice: output buffer size of %d is too small to hold %d encoded samples", dstLen, minLen)
	}
	return nil
}

// checkDecodeBounds does boundary checks for the given buffer sizes.
func checkDecodeBounds(codec Codec, dst []int32, src []byte) error {
	var (
		srcLen = len(src)
		dstLen = len(dst)
		minLen = codec.DecodedLen(srcLen)
	)
	if srcLen == 0 && dstLen == 0 {
		return nil
	} else if minLen == 0 {
		return fmt.Errorf("voice: source buffer size of %d bytes does not align with the block size", srcLen)
	} else if dstLen < minLen {
		return fmt.Errorf("voice: output buffer size of %d is too small to hold %d decoded samples", dstLen, minLen)
	}
	return nil
}

// Closer is the interface that wraps the basic Close method.
type Closer interface {
	Close() error
}

// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) samples into p. It returns the number of samples read
// (0 <= n <= len(p)) and any error encountered. Even if Read returns
// n < len(p), it may use all of p as scratch space during the call. If some
// data is available but not len(p) samples, Read conventionally returns what
// is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after successfully
// reading n > 0 samples, it returns the number of samples read. It may return
// the (non-nil) error from the same call or return the error (and n == 0) from
// a subsequent call. An instance of this general case is that a Reader
// returning a non-zero number of samples at the end of the input stream may
// return either err == EOF or err == nil. The next Read should return 0, EOF.
//
// Callers should always process the n > 0 samples returned before considering
// the error err. Doing so correctly handles I/O errors that happen after
// reading some bytes and also both of the allowed EOF behaviors.
type Reader interface {
	Read(p []int32) (n int, err error)
}

// ReadCloser is the interface that groups the basic Read and Close methods.
type ReadCloser interface {
	Reader
	Closer
}

// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) samples from p to the underlying data stream. It returns
// the number of samples written from p (0 <= n <= len(p)) and any error
// encountered that caused the write to stop early. Write must return a non-nil
// error if it returns n < len(p). Write must not modify the slice data,
// even temporarily.
type Writer interface {
	Write(p []int32) (n int, err error)
}

// WriteCloser is the interface that groups the basic Write and Close methods.
type WriteCloser interface {
	Writer
	Closer
}

func mustDecompress(c []byte) []byte {
	r, err := gzip.NewReader(bytes.NewReader(c))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return b
}
