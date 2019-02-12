package nb

type adpcm struct {
	predictor int
	index     int16
	step      int
	prev      int

	// ms version:
	sample1, sample2 int
	coeff1, coeff2   int
	idelta           int
}

var (
	// CD-ROM XA ADPCM
	adpcmXATable = [5][2]int{
		{0, 0},
		{60, 0},
		{115, -52},
		{98, -55},
		{122, -60},
	}
	adpcmEATable = []int{
		0, 240, 460, 392,
		0, 0, -208, -220,
		0, 1, 3, 4,
		7, 8, 10, 11,
		0, -1, -3, -4,
	}
)
