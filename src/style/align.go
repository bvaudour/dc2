package style

import (
	"unicode/utf8"
	u "util"
)

func slen(s string) int {
	return utf8.RuneCountInString(s)
}

func Left(s string, l int) string {
	if l <= 0 {
		return ""
	}
	d := l - slen(s)
	switch {
	case d > 0:
		return s + u.Repeat(" ", d)
	case d < 0:
		return s[:l]
	default:
		return s
	}
}

func Right(s string, l int) string {
	if l <= 0 {
		return ""
	}
	d := l - slen(s)
	switch {
	case d > 0:
		return u.Repeat(" ", d) + s
	case d < 0:
		return s[-d:]
	default:
		return s
	}
}

func Center(s string, l int) string {
	if l <= 0 {
		return ""
	}
	d := (l - slen(s)) / 2
	switch {
	case d > 0:
		return Left(u.Repeat(" ", d)+s, l)
	case d < 0:
		return s[-d : l-d]
	default:
		return Left(s, l)
	}
}
