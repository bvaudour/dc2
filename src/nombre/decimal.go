package nombre

import (
	"nombre/buffer"
)

type option struct {
	dec int
	fix bool
}

var o = option{dec: 10}

func SetDecMax(d int) {
	if d < 0 {
		o.dec = 10
	} else {
		o.dec = d
	}
}

func SetFixe(f bool) {
	o.fix = f
}

func SetOption(d int, f bool) {
	SetDecMax(d)
	SetFixe(f)
}

func GetOption() (int, bool) {
	return o.dec, o.fix
}

type Decimal struct {
	buf *Entier
	dec int
}

func Int2Decimal(v int) *Decimal {
	out := Decimal{buf: Int2Entier(v)}
	return out.format()
}

func D0() *Decimal {
	return Int2Decimal(0)
}

func D1() *Decimal {
	return Int2Decimal(1)
}

func D10() *Decimal {
	return Int2Decimal(10)
}

func DNone(neg bool) *Decimal {
	return &Decimal{buf: ENone(neg)}
}

func DInf() *Decimal {
	return &Decimal{buf: EInf()}
}

func DNInf() *Decimal {
	return &Decimal{buf: ENInf()}
}

func (n *Decimal) pow() *Entier {
	return Pow10(n.dec)
}

func (n *Decimal) format() *Decimal {
	switch {
	case n.None():
		n.dec = 0
	case o.fix:
		d := o.dec - n.dec
		if d < 0 {
			n.buf = n.buf.By(Pow10(d))
		} else if d > 0 {
			n.buf = n.buf.Time(Pow10(d))
		}
		n.dec = o.dec
	case n.dec > 0:
		for n.dec > 0 {
			div, mod := n.buf.buf.DivMod(10)
			if mod.Is(0) {
				n.buf.buf = div
				n.dec--
			} else {
				break
			}
		}
	case n.dec < 0:
		n.buf = n.buf.Time(n.pow())
		n.dec = 0
	}
	n.buf.format()
	return n
}

func (n *Decimal) IsInt() bool {
	return n.dec == 0
}

func (n *Decimal) Negative() bool {
	return n.buf.Negative()
}

func (n *Decimal) Positive() bool {
	return n.buf.Positive()
}

func (n *Decimal) Even() bool {
	return n.IsInt() && n.buf.Even()
}

func (n *Decimal) Odd() bool {
	return n.IsInt() && n.buf.Odd()
}

func (n *Decimal) Is(v int) bool {
	return n.IsInt() && n.buf.Is(v)
}

func (n *Decimal) None() bool {
	return n.buf.None()
}

func (n *Decimal) Infinite() bool {
	return n.buf.Infinite()
}

func (n *Decimal) NInfinite() bool {
	return n.buf.NInfinite()
}

func (n *Decimal) Copy() *Decimal {
	return &Decimal{n.buf.Copy(), n.dec}
}

func (n *Decimal) SetSign(neg bool) *Decimal {
	n.buf.SetSign(neg)
	return n
}

func (n *Decimal) Abs() *Decimal {
	return &Decimal{n.buf.Abs(), n.dec}
}

func (n *Decimal) Inv() *Decimal {
	return &Decimal{n.buf.Inv(), n.dec}
}

func (n1 *Decimal) compare(n2 *Decimal) int {
	e1, e2 := n1.buf, n2.buf
	d := n1.dec - n2.dec
	switch {
	case d > 0:
		e2 = e2.Time(Pow10(d))
	case d < 0:
		e1 = e1.Time(Pow10(d))
	}
	return e1.compare(e2)
}

func (n1 *Decimal) Compare(n2 Number) int {
	return n1.compare(n2.Decimal())
}

func (n1 *Decimal) Eq(n2 Number) bool {
	return n1.Compare(n2) == 0
}

func (n1 *Decimal) Ne(n2 Number) bool {
	return n1.Compare(n2) != 0
}

func (n1 *Decimal) Ge(n2 Number) bool {
	return n1.Compare(n2) >= 0
}

func (n1 *Decimal) Le(n2 Number) bool {
	return n1.Compare(n2) <= 0
}

func (n1 *Decimal) Gt(n2 Number) bool {
	return n1.Compare(n2) > 0
}

func (n1 *Decimal) Lt(n2 Number) bool {
	return n1.Compare(n2) < 0
}

func (n1 *Decimal) Plus(n2 *Decimal) *Decimal {
	e1, e2 := n1.buf, n2.buf
	d := n1.dec - n2.dec
	dec := n1.dec
	switch {
	case d > 0:
		e2 = e2.Time(Pow10(d))
	case d < 0:
		e1 = e1.Time(Pow10(d))
		dec = n2.dec
	}
	out := Decimal{e1.Plus(e2), dec}
	return out.format()
}

func (n1 *Decimal) Minus(n2 *Decimal) *Decimal {
	return n1.Plus(n2.Inv())
}

func (n1 *Decimal) Time(n2 *Decimal) *Decimal {
	out := new(Decimal)
	out.dec = n1.dec + n2.dec
	out.buf = n1.buf.Time(n2.buf)
	return out.format()
}

func (n1 *Decimal) DivMod(n2 *Decimal) (*Decimal, *Decimal) {
	e1, e2 := n1.buf, n2.buf
	d := n1.dec - n2.dec
	dec := n1.dec
	switch {
	case d > 0:
		e2 = e2.Time(Pow10(d))
	case d < 0:
		e1 = e1.Time(Pow10(d))
		dec = n2.dec
	}
	de, me := e1.DivMod(e2)
	div, mod := Decimal{buf: de}, Decimal{me, dec}
	return div.format(), mod.format()
}

func (n1 *Decimal) By(n2 *Decimal) *Decimal {
	e1, e2 := n1.buf, n2.buf
	d := n1.dec - n2.dec - o.dec
	switch {
	case d > 0:
		e2 = e2.Time(Pow10(d))
	case d < 0:
		e1 = e1.Time(Pow10(d))
	}
	out := Decimal{e1.By(e2), o.dec}
	return out.format()
}

func (n1 *Decimal) Mod(n2 *Decimal) *Decimal {
	_, out := n1.DivMod(n2)
	return out
}

func (n *Decimal) cpw(b *buffer.Buffer) *Decimal {
	switch {
	case b.Is(1):
		return n.Copy()
	case b.Is(2):
		return n.Time(n)
	case b.Even():
		return n.Time(n).cpw(b.By(2))
	default:
		return n.Time(n.Time(n).cpw(b.By(2)))
	}
}
func (n1 *Decimal) Pow(n2 *Decimal) *Decimal {
	switch {
	case !n2.IsInt():
		return DInf()
	case n1.None() || n2.None():
		return DNone(n1.Negative() && n2.Odd())
	case n2.Is(0) || n1.Is(1):
		return D1()
	case n2.Is(1):
		return n1.Copy()
	case n1.Is(0):
		if n2.Negative() {
			return DInf()
		}
		return D0()
	default:
		out := n1.cpw(n2.buf.buf)
		if n2.Negative() {
			return D1().By(out)
		}
		return out
	}
}

func (n *Decimal) Sqrt() *Decimal {
	out := new(Decimal)
	out.dec = o.dec
	e := n.buf
	d := o.dec*2 - n.dec
	switch {
	case d > 0:
		e = e.Time(Pow10(d))
	case d < 0:
		e = e.By(Pow10(d))
	}
	out.buf = e.Sqrt()
	return out.format()
}

func (n *Decimal) String() string {
	out := n.buf.String()
	if n.None() {
		return out
	}
	sign := ""
	if n.Negative() {
		sign = "-"
		out = out[1:]
	}
	for len(out) < n.dec {
		out = "0" + out
	}
	l := len(out) - n.dec
	return sign + out[:l] + "." + out[l:]
}

func (n *Decimal) Sci() string {
	out := n.buf.String()
	if n.None() {
		return out
	}
	sign := ""
	if n.Negative() {
		sign = "-"
		out = out[1:]
	}
	exp := len(out) - 1 - n.dec
	for l := len(out); l > 1 && out[l-1] == '0'; l-- {
		out = out[:l-1]
	}
	return sign + out[0:1] + "." + out[1:] + "E" + int2str(exp)
}
