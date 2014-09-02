package nombre

import (
	"nombre/buffer"
)

type Base struct {
	buf *Entier
	bse int
}

func bse(base int) int {
	switch base {
	case buffer.B2:
		fallthrough
	case buffer.B16:
		return base
	default:
		return buffer.B10
	}
}

func pr(n1, n2 *Base) int {
	switch {
	case n1.bse == buffer.B16 || n2.bse == buffer.B16:
		return buffer.B16
	case n1.bse == buffer.B2 || n2.bse == buffer.B2:
		return buffer.B2
	default:
		return buffer.B10
	}
}

func Int2Base(v, base int) *Base {
	return &Base{Int2Entier(v), bse(base)}
}

func B0(base int) *Base {
	return Int2Base(0, base)
}

func B1(base int) *Base {
	return Int2Base(1, base)
}

func B10(base int) *Base {
	return Int2Base(10, base)
}

func BNone(neg bool, base int) *Base {
	return &Base{ENone(neg), bse(base)}
}

func BInf(base int) *Base {
	return &Base{EInf(), bse(base)}
}

func BNInf(base int) *Base {
	return &Base{ENInf(), bse(base)}
}

func (n *Base) Negative() bool {
	return n.buf.Negative()
}

func (n *Base) Positive() bool {
	return n.buf.Positive()
}

func (n *Base) Even() bool {
	return n.buf.Even()
}

func (n *Base) Odd() bool {
	return n.buf.Odd()
}

func (n *Base) Is(v int) bool {
	return n.buf.Is(v)
}

func (n *Base) None() bool {
	return n.buf.None()
}

func (n *Base) Infinite() bool {
	return n.buf.Infinite()
}

func (n *Base) NInfinite() bool {
	return n.buf.NInfinite()
}

func (n *Base) Copy() *Base {
	return &Base{n.buf.Copy(), n.bse}
}

func (n *Base) SetBase(base int) *Base {
	n.bse = bse(base)
	return n
}

func (n *Base) GetBase() int {
	return n.bse
}

func (n *Base) SetSign(neg bool) *Base {
	n.buf.SetSign(neg)
	return n
}

func (n *Base) Abs() *Base {
	return &Base{n.buf.Abs(), n.bse}
}

func (n *Base) Inv() *Base {
	return &Base{n.buf.Inv(), n.bse}
}

func (n1 *Base) Compare(n2 Number) int {
	return n1.buf.Compare(n2)
}

func (n1 *Base) Eq(n2 Number) bool {
	return n1.Compare(n2) == 0
}

func (n1 *Base) Ne(n2 Number) bool {
	return n1.Compare(n2) != 0
}

func (n1 *Base) Ge(n2 Number) bool {
	return n1.Compare(n2) >= 0
}

func (n1 *Base) Le(n2 Number) bool {
	return n1.Compare(n2) <= 0
}

func (n1 *Base) Gt(n2 Number) bool {
	return n1.Compare(n2) > 0
}

func (n1 *Base) Lt(n2 Number) bool {
	return n1.Compare(n2) < 0
}

func (n1 *Base) Plus(n2 *Base) *Base {
	return &Base{n1.buf.Plus(n2.buf), pr(n1, n2)}
}

func (n1 *Base) Minus(n2 *Base) *Base {
	return &Base{n1.buf.Minus(n2.buf), pr(n1, n2)}
}

func (n1 *Base) Time(n2 *Base) *Base {
	return &Base{n1.buf.Time(n2.buf), pr(n1, n2)}
}

func (n1 *Base) DivMod(n2 *Base) (*Base, *Base) {
	div, mod := n1.buf.DivMod(n2.buf)
	bse := pr(n1, n2)
	return &Base{div, bse}, &Base{mod, bse}
}

func (n1 *Base) By(n2 *Base) *Base {
	return &Base{n1.buf.By(n2.buf), pr(n1, n2)}
}

func (n1 *Base) Mod(n2 *Base) *Base {
	return &Base{n1.buf.Mod(n2.buf), pr(n1, n2)}
}

func (n1 *Base) Pow(n2 *Base) *Base {
	return &Base{n1.buf.Pow(n2.buf), pr(n1, n2)}
}

func (n *Base) Sqrt() *Base {
	return &Base{n.buf.Sqrt(), n.bse}
}

func (n *Base) Fact() *Base {
	return &Base{n.buf.Fact(), n.bse}
}

func (n *Base) String() string {
	return n.buf.StrBase(n.bse)
}

func (n *Base) Sci() string {
	return n.String()
}
