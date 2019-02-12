package vocoder

// IsAvailable returns true if dv is built with cgo support.
func IsAvailable() bool {
	return isAvailable()
}
