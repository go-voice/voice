/*
Package codec2 implements the Codec2 speech codec
*/
package codec2

// IsAvailable returns true if voice is built with codec2 support.
func IsAvailable() bool {
	return isAvailable()
}

// Mode setting
type Mode int

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
	Mode3200 Mode = iota
	Mode2400
	Mode1600
	Mode1400
	Mode1300
	Mode1200
	Mode700
	Mode700B
	Mode700C
)
