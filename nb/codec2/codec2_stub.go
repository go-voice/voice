// +build !cgo

package codec2

import (
	"errors"

	"github.com/go-voice/voice/nb"
)

func isAvailable() bool {
	return false
}

func New(mode Mode) (nb.CodecCloser, error) {
	return nil, errors.New("codec2: not avaiable, build this package with cgo support")
}
