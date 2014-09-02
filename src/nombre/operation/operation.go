package operation

import "math"

const (
	B10E9 uint = 1000000000
	B2E32 uint = 1 << 32
	MSK32 uint = B2E32 - 1
)

func Compare(i1, i2 uint) int {
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

func Concat(i1 uint, i2 uint) uint {
	return (i1 << 32) | i2
}

func Split(i uint) (uint, uint) {
	return i >> 32, i & MSK32
}

func Result(i uint) uint {
	_, out := Split(i)
	return out
}

func Remainder(i uint) uint {
	out, _ := Split(i)
	return out
}

func BitLen(i uint) uint {
	var out uint
	for {
		if i == 0 {
			break
		}
		i >>= 1
		out++
	}
	return out
}

func LeftShift(i1, i2, s uint) (uint, uint) {
	return Split((i1 << s) | i2)
}

func RightShift(i1, i2, s uint) (uint, uint) {
	return (i1 | (i2 >> s)), (i2 << (32 - s) & MSK32)
}

func Sqrt(i uint) uint {
	return uint(math.Sqrt(float64(i)))
}
