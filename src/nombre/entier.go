package nombre

import (
	"nombre/buffer"
	"strconv"
)

type Entier struct {
	buf *buffer.Buffer
	neg bool
}

func Int2Entier(v int) *Entier {
	return &Entier{buffer.NewBuffer(int2uint(v)), v < 0}
}

func Bool2Entier(v bool) *Entier {
	if v {
		return E1()
	}
	return E0()
}

func E0() *Entier {
	return Int2Entier(0)
}

func E1() *Entier {
	return Int2Entier(1)
}

func E10() *Entier {
	return Int2Entier(10)
}

func ENone(neg bool) *Entier {
	return &Entier{buffer.EmptyBuffer(), neg}
}

func EInf() *Entier {
	return ENone(false)
}

func ENInf() *Entier {
	return ENone(true)
}

func Pow10(v int) *Entier {
	return &Entier{buf: buffer.Pow10(int2uint(v))}
}

func (n *Entier) format() *Entier {
	if n.neg && n.buf.Is(0) {
		n.neg = false
	}
	return n
}

func (n *Entier) Negative() bool {
	return n.neg
}

func (n *Entier) Positive() bool {
	return !n.neg
}

func (n *Entier) Even() bool {
	return n.buf.Even()
}

func (n *Entier) Odd() bool {
	return n.buf.Odd()
}

func (n *Entier) Is(v int) bool {
	return v < 0 == n.Negative() && n.buf.Is(int2uint(v))
}

func (n *Entier) None() bool {
	return n.buf.Empty()
}

func (n *Entier) Infinite() bool {
	return n.None() && n.Positive()
}

func (n *Entier) NInfinite() bool {
	return n.None() && n.Negative()
}

func (n *Entier) Copy() *Entier {
	return &Entier{n.buf.Copy(), n.neg}
}

func (n *Entier) SetSign(neg bool) *Entier {
	n.neg = neg
	return n.format()
}

func (n *Entier) Abs() *Entier {
	return n.Copy().SetSign(false)
}

func (n *Entier) Inv() *Entier {
	return n.Copy().SetSign(!n.neg)
}

func (n1 *Entier) compare(n2 *Entier) int {
	switch {
	case n1.Positive():
		if n2.Positive() {
			return n1.buf.Compare(n2.buf)
		}
		return 1
	case n2.Positive():
		return -1
	default:
		return n2.buf.Compare(n1.buf)
	}
}

func (n1 *Entier) Compare(n2 Number) int {
	if n2.IsDecimal() {
		return n1.Decimal().compare(n2.Decimal())
	}
	return n1.compare(n2.Entier())
}

func (n1 *Entier) Eq(n2 Number) bool {
	return n1.Compare(n2) == 0
}

func (n1 *Entier) Ne(n2 Number) bool {
	return n1.Compare(n2) != 0
}

func (n1 *Entier) Ge(n2 Number) bool {
	return n1.Compare(n2) >= 0
}

func (n1 *Entier) Le(n2 Number) bool {
	return n1.Compare(n2) <= 0
}

func (n1 *Entier) Gt(n2 Number) bool {
	return n1.Compare(n2) > 0
}

func (n1 *Entier) Lt(n2 Number) bool {
	return n1.Compare(n2) < 0
}

func (n1 *Entier) Plus(n2 *Entier) *Entier {
	out := new(Entier)
	switch {
	case n1.neg == n2.neg:
		out.buf = buffer.Plus(n1.buf, n2.buf)
		out.neg = n1.neg
	case buffer.Compare(n1.buf, n2.buf) >= 0:
		out.buf = buffer.Minus(n1.buf, n2.buf)
		out.neg = n1.neg
	default:
		out.buf = buffer.Minus(n2.buf, n1.buf)
		out.neg = n2.neg
	}
	return out.format()
}

func (n1 *Entier) Minus(n2 *Entier) *Entier {
	return n1.Plus(n2.Inv())
}

func (n1 *Entier) Time(n2 *Entier) *Entier {
	out := new(Entier)
	out.neg = n1.neg != n2.neg
	out.buf = buffer.Time(n1.buf, n2.buf)
	return out.format()
}

func (n1 *Entier) DivMod(n2 *Entier) (*Entier, *Entier) {
	div, mod := new(Entier), new(Entier)
	div.neg, mod.neg = n1.neg != n2.neg, n1.neg
	div.buf, mod.buf = buffer.DivMod(n1.buf, n2.buf)
	return div.format(), mod.format()
}

func (n1 *Entier) By(n2 *Entier) *Entier {
	out, _ := n1.DivMod(n2)
	return out
}

func (n1 *Entier) Mod(n2 *Entier) *Entier {
	_, out := n1.DivMod(n2)
	return out
}

func (n1 *Entier) Pow(n2 *Entier) *Entier {
	out := new(Entier)
	if n1.Negative() && n2.Odd() {
		out.neg = true
	}
	switch {
	case n1.None() || n2.None():
		out.buf = buffer.EmptyBuffer()
	case n2.Negative():
		if n1.buf.Is(0) {
			out.buf = buffer.EmptyBuffer()
		} else {
			out.buf = buffer.NewBuffer(0)
		}
	default:
		out.buf = buffer.Pow(n1.buf, n2.buf)
	}
	return out.format()
}

func (n *Entier) Sqrt() *Entier {
	if n.Negative() {
		return EInf()
	}
	return &Entier{buf: n.buf.Sqrt()}
}

func (n *Entier) Fact() *Entier {
	if n.Negative() {
		return EInf()
	}
	return &Entier{buf: n.buf.Fact()}
}

func (n *Entier) String() string {
	return n.StrBase(buffer.B10)
}

func (n *Entier) Sci() string {
	out := n.String()
	if n.None() {
		return out
	}
	dec := 1
	if n.Negative() {
		dec++
	}
	exp := strconv.Itoa(len(out) - dec)
	if dec != len(out) {
		out = out[:dec] + "." + out[dec:]
	}
	return out + "E" + exp
}

func (n *Entier) StrBase(bse int) string {
	out := ""
	switch {
	case bse == buffer.B2:
		if n.Negative() {
			out += "1"
		} else {
			out += "0"
		}
		out += "%"
	case bse == buffer.B16:
		if n.Negative() {
			out += "1"
		} else {
			out += "0"
		}
		out += "x"
	default:
		if n.Negative() {
			out += "-"
		}
	}
	out += buffer.Format(n.buf, bse)
	return out
}
