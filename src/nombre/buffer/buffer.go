package buffer

import (
	u "util"
)

type Buffer []uint

func BMake(sze int) Buffer {
	return Buffer(make([]uint, sze))
}

func BAppend(b Buffer, args ...uint) Buffer {
	return Buffer(append([]uint(b), args...))
}

func (b *Buffer) Len() int {
	return len(*b)
}

func (b *Buffer) Size() uint {
	return uint(b.Len())
}

func (b *Buffer) Empty() bool {
	return b.Len() == 0
}

func (b *Buffer) get0(idx int) uint {
	return (*b)[idx]
}

func (b *Buffer) set0(idx int, v uint) {
	(*b)[idx] = v
}

func (b *Buffer) from0(f int) *Buffer {
	out := (*b)[f:]
	return &out
}

func (b *Buffer) to0(t int) *Buffer {
	out := (*b)[:t]
	return &out
}

func (b *Buffer) slice0(f, t int) *Buffer {
	out := (*b)[f:t]
	return &out
}

func (b *Buffer) Get(idx int) uint {
	return b.get0(index(idx, b))
}

func (b *Buffer) Set(idx int, v uint) {
	b.set0(index(idx, b), v)
}

func (b *Buffer) From(fr int) *Buffer {
	return b.from0(slice(fr, b))
}

func (b *Buffer) To(to int) *Buffer {
	return b.to0(slice(to, b))
}

func (b *Buffer) Slice(fr, to int) *Buffer {
	return b.slice0(slice(fr, b), slice(to, b))
}

func MakeBuffer(sze int) *Buffer {
	out := BMake(sze)
	return &out
}

func NewBuffer(args ...uint) *Buffer {
	out := BMake(len(args))
	for i, a := range args {
		out[i] = a
	}
	return &out
}

func EmptyBuffer() *Buffer {
	return MakeBuffer(0)
}

func (b *Buffer) ToArray() []uint {
	return []uint(*b)
}

func (b *Buffer) Add(args ...uint) *Buffer {
	*b = BAppend(*b, args...)
	return b
}

func (b *Buffer) Insert(args ...uint) *Buffer {
	*b = Buffer(append(args, *b...))
	return b
}

func (b *Buffer) Clear() *Buffer {
	*b = BMake(0)
	return b
}

func (b *Buffer) It(idx int) Itr {
	return &bit{b, index(idx, b)}
}

func (b *Buffer) CountZeros() int {
	out := 0
	for it := b.It(0); it.Exists() && it.Get() == 0; it.Next() {
		out++
	}
	return out
}

func (b *Buffer) Format() *Buffer {
	if b.Empty() {
		return b
	}
	out := b.from0(b.To(-1).CountZeros())
	*b = *out
	return b
}

func (b *Buffer) Trim() *Buffer {
	out := b.from0(b.CountZeros())
	*b = *out
	return b
}

func (b Buffer) Copy() *Buffer {
	return NewBuffer(b...)
}

func (b *Buffer) AddZeros(z int) *Buffer {
	return b.Add(make([]uint, z)...)
}

func (b *Buffer) Mkstring(begin, separator, end string) string {
	if b.Empty() {
		return begin + end
	}
	it := b.It(0)
	out := begin + it.String()
	for it.Next() {
		out += (separator + it.String())
	}
	return out + end
}

func (b *Buffer) String() string {
	//return b.Mkstring("[", "; ", "]")
	return u.SliceString(*b)
}

func (b *Buffer) Is(v uint) bool {
	if b.Empty() {
		return false
	}
	if b.Get(-1) != v {
		return false
	}
	sl := b.To(-1)
	return sl.CountZeros() == sl.Len()
}

func (b *Buffer) Even() bool {
	return b.Empty() || b.Get(-1)%2 == 0
}

func (b *Buffer) Odd() bool {
	return !b.Even()
}

func (b *Buffer) ToInt() int {
	if b.Len() != 1 {
		return -1
	}
	return int(b.get0(0))
}
