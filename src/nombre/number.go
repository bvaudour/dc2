package nombre

import (
	"nombre/buffer"
)

type Number interface {
	IsEntier() bool
	IsBase() bool
	IsDecimal() bool
	Entier() *Entier
	Base() *Base
	Decimal() *Decimal
	String() string
	Sci() string
	Negative() bool
	Positive() bool
	Even() bool
	Odd() bool
	Is(int) bool
	None() bool
	Infinite() bool
	NInfinite() bool
	Compare(n Number) int
	Eq(n Number) bool
	Ne(n Number) bool
	Ge(n Number) bool
	Le(n Number) bool
	Gt(n Number) bool
	Lt(n Number) bool
}

func (n *Entier) IsEntier() bool {
	return true
}

func (n *Entier) IsBase() bool {
	return false
}

func (n *Entier) IsDecimal() bool {
	return false
}

func (n *Entier) Entier() *Entier {
	return n
}

func (n *Entier) Base() *Base {
	return &Base{n, buffer.B10}
}

func (n *Entier) Decimal() *Decimal {
	out := &Decimal{buf: n.Copy()}
	return out.format()
}

func (n *Base) IsEntier() bool {
	return false
}

func (n *Base) IsBase() bool {
	return true
}

func (n *Base) IsDecimal() bool {
	return false
}

func (n *Base) Entier() *Entier {
	return n.buf
}

func (n *Base) Base() *Base {
	return n
}

func (n *Base) Decimal() *Decimal {
	return n.Entier().Decimal()
}

func (n *Decimal) IsEntier() bool {
	return false
}

func (n *Decimal) IsBase() bool {
	return false
}

func (n *Decimal) IsDecimal() bool {
	return true
}

func (n *Decimal) Entier() *Entier {
	return n.buf.By(n.pow())
}

func (n *Decimal) Base() *Base {
	return n.Entier().Base()
}

func (n *Decimal) Decimal() *Decimal {
	return n
}
