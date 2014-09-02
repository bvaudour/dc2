package buffer

func index0(idx, max int) int {
	switch {
	case idx >= 0 && idx < max:
		return idx
	case idx < 0 && idx >= -max:
		return max + idx
	default:
		return -1
	}
}

func index1(idx, max int) int {
	if idx < 0 {
		return max + idx
	}
	return idx
}

func index(idx int, b *Buffer) int {
	return index0(idx, b.Len())
}

func slice(sl int, b *Buffer) int {
	return index1(sl, b.Len())
}

func Valid(idx int, b *Buffer) bool {
	return idx >= 0 && idx < b.Len()
}
