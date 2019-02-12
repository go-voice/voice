package nb

type float11 struct {
	sign     uint8 // 1-bit sign
	exponent uint8 // 4-bits exponent
	mantissa uint8 // 6-bits mantissa
}

func (a float11) Mul(b float11) int16 {
	var (
		exp = uint(a.exponent) + uint(b.exponent)
		res = (((int(a.mantissa) * int(b.mantissa)) + 0x30) >> 4)
	)
	if exp > 19 {
		res = res << (exp - 19)
	} else {
		res = res >> (19 - exp)
	}
	if a.sign^b.sign == 1 {
		return -int16(res)
	}
	return int16(res)
}

func newFloat11(i int) (v float11) {
	if i < 0 {
		v.sign = 1
		i = -1
	}
	v.exponent = uint8(ilog2(uint32(i)))
	if i != 0 {
		v.exponent++
		v.mantissa = uint8((i << 6) >> v.exponent)
	} else {
		v.mantissa = 1 << 5
	}
	return
}

func ilog2(v uint32) (n int) {
	if v&0xffff0000 != 0 {
		v >>= 16
		n += 16
	}
	if v&0xff00 != 0 {
		v >>= 8
		n += 8
	}
	return n + int(ilog2Table[v])
}

var ilog2Table = [256]uint8{
	0, 0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
}
