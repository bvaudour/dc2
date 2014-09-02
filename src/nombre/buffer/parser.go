package buffer

import (
	op "nombre/operation"
	"strconv"
	"strings"
)

const (
	B2  int    = 2
	B10 int    = 10
	B16 int    = 16
	INF string = "<infinite>"
)

func int2String(v uint, base int) string {
	return strings.ToUpper(strconv.FormatUint(uint64(v), base))
}

func string2Int(s string, base int) uint {
	v, _ := strconv.ParseUint(s, base, 32)
	return uint(v)
}

func formatString(s string, length int) string {
	out := s
	for len(out) < length {
		out = "0" + out
	}
	return out
}

func string2Buffer(s string, base int, bits int) *Buffer {
	if s == "" {
		return EmptyBuffer()
	}
	l := len(s)
	i := l % bits
	var out *Buffer
	var it Itr
	if i != 0 {
		out = MakeBuffer(l/bits + 1)
		out.Set(0, string2Int(s[:i], base))
		it = out.It(1)
	} else {
		out = MakeBuffer(l / bits)
		it = out.It(0)
	}
	for ; i < l; i += bits {
		it.Set(string2Int(s[i:i+bits], base))
		it.Next()
	}
	return out.Format()
}

func buffer2String(b *Buffer, base int, bits int) string {
	if b.Empty() {
		return INF
	}
	it := b.It(0)
	out := int2String(it.Get(), base)
	for it.Next() {
		out += formatString(int2String(it.Get(), base), bits)
	}
	return out
}

func Parse(s string, base int) *Buffer {
	switch base {
	case B2:
		return string2Buffer(s, base, 32)
	case B10:
		return string2Buffer(s, base, 9).Convert(op.B10E9, op.B2E32)
	case B16:
		return string2Buffer(s, base, 8)
	default:
		return EmptyBuffer()
	}
}

func Format(b *Buffer, base int) string {
	switch base {
	case B2:
		return buffer2String(b, base, 32)
	case B10:
		return buffer2String(b.Convert(op.B2E32, op.B10E9), base, 9)
	case B16:
		return buffer2String(b, base, 8)
	default:
		return INF
	}
}
