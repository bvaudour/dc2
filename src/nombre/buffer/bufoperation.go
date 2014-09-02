package buffer

import (
	op "nombre/operation"
)

func (b *Buffer) BitLen() uint {
	return op.BitLen(b.get0(0))
}

func (b *Buffer) LeftShift(n uint) *Buffer {
	if b.Empty() {
		return b
	}
	var v, r uint
	for it := b.It(-1); it.Exists(); it.Prev() {
		r, v = op.LeftShift(it.Get(), r, n)
		it.Set(v)
	}
	b.Insert(r)
	return b.Format()
}

func (b *Buffer) RightShift(n uint) *Buffer {
	if b.Empty() {
		return b
	}
	var v, r uint
	for it := b.It(0); it.Exists(); it.Next() {
		v, r = op.RightShift(r, it.Get(), n)
		it.Set(v)
	}
	return b.Format()
}

func (b *Buffer) Convert(oldBase, newBase uint) *Buffer {
	if b.Empty() || oldBase < 2 || newBase < 2 {
		return EmptyBuffer()
	}
	c := b.Copy()
	if oldBase == newBase {
		return c
	}
	out := EmptyBuffer()
	for !c.Empty() {
		var r uint = 0
		for it := c.It(0); it.Exists(); it.Next() {
			o := r*oldBase + it.Get()
			it.Set(o / newBase)
			r = o % newBase
		}
		out.Insert(r)
		c.Trim()
	}
	return out.Format()
}

func (b *Buffer) Plus(v uint) *Buffer {
	if b.Empty() {
		return EmptyBuffer()
	}
	out := MakeBuffer(b.Len() + 1)
	var r uint = v
	itO := out.It(-1)
	for itI := b.It(-1); itI.Exists(); itI.Prev() {
		c := itI.Get() + r
		itO.Set(op.Result(c))
		r = op.Remainder(c)
		itO.Prev()
	}
	itO.Set(r)
	return out.Format()
}

func (b *Buffer) Minus(v uint) *Buffer {
	if b.Empty() {
		return EmptyBuffer()
	}
	out := MakeBuffer(b.Len())
	itO := out.It(-1)
	var r uint = v
	for itI := b.It(-1); itI.Exists(); itI.Prev() {
		c := itI.Get() - r
		itO.Set(op.Result(c))
		if op.Remainder(c) == 0 {
			r = 0
		} else {
			r = 1
		}
		itO.Prev()
	}
	if r != 0 {
		return EmptyBuffer()
	}
	return out.Format()
}

func (b *Buffer) Time(v uint) *Buffer {
	if b.Empty() {
		return EmptyBuffer()
	}
	switch v {
	case 0:
		return NewBuffer(0)
	case 1:
		return b.Copy().Format()
	}
	out := MakeBuffer(b.Len() + 1)
	itO := out.It(-1)
	var r uint = 0
	for itI := b.It(-1); itI.Exists(); itI.Prev() {
		c := itI.Get()*v + r
		itO.Set(op.Result(c))
		r = op.Remainder(c)
		itO.Prev()
	}
	itO.Set(r)
	return out.Format()
}

func (b *Buffer) DivMod(v uint) (*Buffer, *Buffer) {
	if b.Empty() || v == 0 {
		return EmptyBuffer(), EmptyBuffer()
	}
	var r uint = 0
	div := MakeBuffer(b.Len())
	itO := div.It(0)
	for itI := b.It(0); itI.Exists(); itI.Next() {
		c := op.Concat(r, itI.Get())
		itO.Set(c / v)
		r = c % v
		itO.Next()
	}
	return div.Format(), NewBuffer(r)
}

func (b *Buffer) By(v uint) *Buffer {
	out, _ := b.DivMod(v)
	return out
}

func (b *Buffer) Mod(v uint) *Buffer {
	_, out := b.DivMod(v)
	return out
}

func (b *Buffer) Pow(v uint) *Buffer {
	switch {
	case b.Empty():
		return EmptyBuffer()
	case v == 0:
		return NewBuffer(1)
	case v == 1 || b.Is(0) || b.Is(1):
		return b.Copy()
	case v == 2:
		return Time(b, b)
	default:
		n := Time(b, b).Pow(v >> 1)
		if v%2 == 0 {
			return n
		} else {
			return Time(b, n)
		}
	}
}

func (b1 *Buffer) Compare(b2 *Buffer) int {
	return Compare(b1, b2)
}

func Compare(b1, b2 *Buffer) int {
	if b1.Empty() {
		if b2.Empty() {
			return 0
		}
		return 1
	}
	if b2.Empty() {
		return -1
	}
	if c := op.Compare(b1.Size(), b2.Size()); c != 0 {
		return c
	}
	for it1, it2 := b1.It(0), b2.It(0); it1.Exists(); it1.Next() {
		if c := op.Compare(it1.Get(), it2.Get()); c != 0 {
			return c
		}
		it2.Next()
	}
	return 0
}

func Increment(b *Buffer) *Buffer {
	return b.Plus(1)
}

func Decrement(b *Buffer) *Buffer {
	return b.Minus(1)
}

func Pow10(p uint) *Buffer {
	return NewBuffer(10).Pow(p)
}

func Plus(b1, b2 *Buffer) *Buffer {
	if b1.Empty() || b2.Empty() {
		return EmptyBuffer()
	}
	l := b1.Len()
	if b2.Len() > l {
		l = b2.Len()
	}
	out := MakeBuffer(l + 1)
	itO := out.It(-1)
	it1, it2 := b1.It(-1), b2.It(-1)
	var r uint = 0
	for ; it1.Exists() && it2.Exists(); itO.Prev() {
		c := r + it1.Get() + it2.Get()
		itO.Set(op.Result(c))
		r = op.Remainder(c)
		it1.Prev()
		it2.Prev()
	}
	for ; it1.Exists(); itO.Prev() {
		c := r + it1.Get()
		itO.Set(op.Result(c))
		r = op.Remainder(c)
		it1.Prev()
	}
	for ; it2.Exists(); itO.Prev() {
		c := r + it2.Get()
		itO.Set(op.Result(c))
		r = op.Remainder(c)
		it2.Prev()
	}
	itO.Set(r)
	return out.Format()
}

func Minus(b1, b2 *Buffer) *Buffer {
	if b1.Empty() || b2.Empty() || Compare(b1, b2) < 0 {
		return EmptyBuffer()
	}
	out := MakeBuffer(b1.Len())
	itO := out.It(-1)
	it1, it2 := b1.It(-1), b2.It(-1)
	var r uint = 0
	for ; it1.Exists() && it2.Exists(); itO.Prev() {
		c := it1.Get() - it2.Get() - r
		itO.Set(op.Result(c))
		if op.Remainder(c) == 0 {
			r = 0
		} else {
			r = 1
		}
		it1.Prev()
		it2.Prev()
	}
	for ; it1.Exists(); itO.Prev() {
		c := it1.Get() - r
		itO.Set(op.Result(c))
		if op.Remainder(c) == 0 {
			r = 0
		} else {
			r = 1
		}
		it1.Prev()
	}
	return out.Format()
}

func Time(b1, b2 *Buffer) *Buffer {
	if b1.Empty() || b2.Empty() {
		return EmptyBuffer()
	}
	out := NewBuffer(0)
	z := 0
	for it := b2.It(-1); it.Exists(); it.Prev() {
		r := b1.Time(it.Get())
		out = Plus(out, r.AddZeros(z))
		z++
	}
	return out
}

func setDiv(b1, b2 *Buffer) uint {
	div := op.Concat(b1.Get(0), b1.Get(1)) / b2.Get(0)
	if div > op.MSK32 {
		return op.MSK32
	}
	m := b2.Time(div)
	if Compare(m, b1) > 0 {
		m = Minus(m, b2)
		div--
	}
	m = Minus(b1, m)
	it1 := b1.It(-1)
	for it2 := m.It(-1); it2.Exists(); it2.Prev() {
		it1.Set(it2.Get())
		it1.Prev()
	}
	for ; it1.Exists(); it1.Prev() {
		it1.Set(0)
	}
	return div
}

func DivMod(b1, b2 *Buffer) (*Buffer, *Buffer) {
	if b1.Empty() || b2.Empty() || b2.Is(0) {
		return EmptyBuffer(), EmptyBuffer()
	}
	switch Compare(b1, b2) {
	case 0:
		return NewBuffer(1), NewBuffer(0)
	case -1:
		return NewBuffer(0), b1.Copy()
	}
	if b2.Len() == 1 {
		return b1.DivMod(b2.get0(0))
	}
	bitl := 32 - b2.BitLen()
	k1, k2 := b1.Copy().LeftShift(bitl), b2.Copy().LeftShift(bitl)
	if k1.Len() == b1.Len() {
		k1.Insert(0)
	}
	l := k2.Len() + 1
	div := MakeBuffer(k1.Len() - k2.Len())
	for i := 0; i < div.Len(); i++ {
		div.set0(i, setDiv(k1.Slice(i, i+l), k2))
	}
	return div.Format(), k1.RightShift(bitl)
}

func By(b1, b2 *Buffer) *Buffer {
	out, _ := DivMod(b1, b2)
	return out
}

func Mod(b1, b2 *Buffer) *Buffer {
	_, out := DivMod(b1, b2)
	return out
}

func Pow(b1, b2 *Buffer) *Buffer {
	switch {
	case b1.Empty() || b2.Empty():
		return EmptyBuffer()
	case b2.Is(0):
		return NewBuffer(1)
	case b1.Is(0) || b1.Is(1) || b2.Is(1):
		return b1.Copy()
	case b2.Is(2):
		return Time(b1, b1)
	default:
		n := Pow(Time(b1, b1), b2.Copy().RightShift(1))
		if b2.Even() {
			return n
		}
		return Time(b1, n)
	}
}

func heron(root, orig *Buffer) (*Buffer, bool) {
	out := Plus(root, By(orig, root)).RightShift(1)
	ok := Compare(out, root) == 0
	return out, ok
}

func (b *Buffer) Sqrt() *Buffer {
	switch b.Len() {
	case 0:
		return EmptyBuffer()
	case 1:
		return NewBuffer(op.Sqrt(b.Get(0)))
	case 2:
		return NewBuffer(op.Sqrt(op.Concat(b.Get(0), b.Get(1))))
	}
	var out *Buffer
	var itO Itr
	itI := b.It(0)
	if b.Len()%2 == 0 {
		out = MakeBuffer(b.Len() >> 1)
		itO = out.It(0)
	} else {
		out = MakeBuffer((b.Len() >> 1) + 1)
		itO = out.It(0)
		itO.Set(op.Sqrt(itI.Get()))
		itO.Next()
		itI.Next()
	}
	for ; itI.Exists(); itI.Next() {
		v := itI.Get()
		itI.Next()
		itO.Set(op.Sqrt(op.Concat(v, itI.Get())))
		itO.Next()
	}
	out = out.Format()
	for ok := false; !ok; {
		out, ok = heron(out, b)
	}
	return out
}

func (b *Buffer) Fact() *Buffer {
	if b.Empty() {
		return EmptyBuffer()
	}
	one := NewBuffer(1)
	if Compare(b, one) <= 0 {
		return one
	}
	out := b.Copy()
	for elt := Decrement(out); Compare(elt, one) > 0; elt = Decrement(elt) {
		out = Time(out, elt)
	}
	return out
}
