package voice

type Tag uint16

// Tag bases
const (
	PCMBase    Tag = 0x0100
	Codec2Base Tag = 0x7300
	SpeexBase  Tag = 0x7340
	MBEBase    Tag = 0x7380
)

// PCM formats
const (
	PCMS8    Tag = PCMBase + iota //   signed  8-bit
	PCMU8                         // unsigned  8-bit
	PCMULAW                       //    A-Law  8-bit
	PCMALAW                       //    µ-Law  8-bit
	PCMS16LE                      //   signed 16-bit little-endian
	PCMS16BE                      //   signed 16-bit big-endian
	PCMU16LE                      // unsigned 16-bit little-endian
	PCMU16BE                      // unsigned 16-bit big-endian
	PCMS24LE                      //   signed 24-bit little-endian
	PCMS24BE                      //   signed 24-bit    big-endian
	PCMU24LE                      // unsigned 24-bit little-endian
	PCMU24BE                      // unsigned 24-bit    big-endian
	PCMS32LE                      //   signed 32-bit little-endian
	PCMS32BE                      //   signed 32-bit    big-endian
	PCMU32LE                      // unsigned 32-bit little-endian
	PCMU32BE                      // unsigned 32-bit    big-endian
	PCMS64LE                      //   signed 64-bit little-endian
	PCMS64BE                      //   signed 64-bit    big-endian
	PCMF32BE                      //    float 32-bit    big-endian
	PCMF32LE                      //    float 32-bit little-endian
	PCMF64BE                      //    float 64-bit    big-endian
	PCMF64LE                      //    float 64-bit little-endian
)

// Codec2 formats
const (
	Codec2Mode3200 Tag = Codec2Base + iota // Codec2 3200 bps
	Codec2Mode2400                         // Codec2 2400 bps
	Codec2Mode1600                         // Codec2 1600 bps
	Codec2Mode1400                         // Codec2 1400 bps
	Codec2Mode1200                         // Codec2 1200 bps
	Codec2Mode700                          // Codec2  700 bps
	Codec2Mode700B                         // Codec2  700 bps variant B
	Codec2Mode700C                         // Codec2  700 bps variant C
	Codec2ModeWB                           // Codec2 wideband
)

// MBE formats
const (
	AMBE3600x2400 Tag = MBEBase + iota // AMBE 3600 x 2400 bps
	AMBE3600x2450                      // AMBE 3600 x 2450 bps
	IMBE7100x4400                      // IMBE 7100 x 4400 bps
	IMBE7200x4400                      // IMBE 7200 x 4400 bps
)
