package nombre

import (
	"nombre/buffer"
	"regexp"
	"strconv"
	"strings"
)

const (
	entier = iota
	base2
	base16
	decimal
	scient
	unknown
)

var reg map[uint64]*regexp.Regexp = initR()

func initR() map[uint64]*regexp.Regexp {
	out := make(map[uint64]*regexp.Regexp)
	out[entier] = regexp.MustCompile(`^(\+|-)?\d+$`)
	out[base2] = regexp.MustCompile(`^(0|1)%(0|1)+$`)
	out[base16] = regexp.MustCompile(`^(0|1)(X|x)[0-9A-Fa-f]+$`)
	out[decimal] = regexp.MustCompile(`^(\+|-)?(\d+\.\d*|\.\d+)$`)
	out[scient] = regexp.MustCompile(`^(\+|-)?(\d+(\.\d*)?|\.\d+)?(E|e)(\+|-)?\d+$`)
	return out
}

func int2uint(v int) uint {
	if v < 0 {
		return uint(-v)
	}
	return uint(v)
}

func int2str(v int) string {
	return strconv.Itoa(v)
}

func str2int(s string) int {
	out, _ := strconv.Atoi(s)
	return out
}

func match(reg, s string) bool {
	ok, _ := regexp.MatchString(reg, s)
	return ok
}

func Analyse(s string) uint64 {
	for a, r := range reg {
		if r.MatchString(s) {
			return a
		}
	}
	return unknown
}

func IsNumber(s string) bool {
	return Analyse(s) != unknown
}

func parse10(s string) *Entier {
	out := new(Entier)
	switch {
	case s[0] == '-':
		out.neg = true
		fallthrough
	case s[0] == '+':
		out.buf = buffer.Parse(s[1:], 10)
	default:
		out.buf = buffer.Parse(s, 10)
	}
	return out.format()
}

func parseBase(s string, base int) *Base {
	e := new(Entier)
	e.neg = s[0] == '1'
	e.buf = buffer.Parse(s[2:], base)
	return &Base{e.format(), base}
}

func parse2(s string) *Base {
	return parseBase(s, buffer.B2)
}

func parse16(s string) *Base {
	return parseBase(s, buffer.B16)
}

func parseDec(s string) *Decimal {
	out := new(Decimal)
	d := strings.Index(s, ".")
	out.dec = len(s) - 1 - d
	out.buf = parse10(s[:d] + s[d+1:])
	return out.format()
}

func parseSci(s string) *Decimal {
	out := new(Decimal)
	pE := strings.Index(strings.ToLower(s), "e")
	dc, exp := s[:pE], s[pE+1:]
	switch {
	case dc == "" || dc == "+":
		out.buf = E1()
	case dc == "-":
		out.buf = Int2Entier(-1)
	case strings.Contains(dc, "."):
		out = parseDec(dc)
	default:
		out.buf = parse10(dc)
	}
	out.dec -= str2int(exp)
	return out.format()
}

func Parse(s string) Number {
	switch Analyse(s) {
	case entier:
		return parse10(s)
	case decimal:
		return parseDec(s)
	case scient:
		return parseSci(s)
	case base2:
		return parse2(s)
	case base16:
		return parse16(s)
	default:
		return EInf()
	}
}

func ParseEntier(s string) *Entier {
	return Parse(s).Entier()
}

func ParseDecimal(s string) *Decimal {
	return Parse(s).Decimal()
}

func ParseBase(s string) *Base {
	return Parse(s).Base()
}
