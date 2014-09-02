package buffer

type Itr interface {
	End() bool
	Exists() bool
	First() bool
	Last() bool
	Next() bool
	Prev() bool
	Jump(int) bool
	Set(uint) bool
	Get() uint
	Index() int
	String() string
}

type bit struct {
	buf *Buffer
	idx int
}

func (it *bit) End() bool {
	return it.Index() < 0
}

func (it *bit) Exists() bool {
	return !it.End()
}

func (it *bit) First() bool {
	it.idx = index(0, it.buf)
	return it.Exists()
}

func (it *bit) Last() bool {
	it.idx = index(-1, it.buf)
	return it.Exists()
}

func (it *bit) Next() bool {
	return it.Jump(1)
}

func (it *bit) Prev() bool {
	return it.Jump(-1)
}

func (it *bit) Jump(incr int) bool {
	if it.End() {
		return false
	}
	idx := it.Index() + incr
	if Valid(idx, it.buf) {
		it.idx = idx
	} else {
		it.idx = -1
	}
	return it.Exists()
}

func (it *bit) Set(v uint) bool {
	if it.End() {
		return false
	}
	it.buf.set0(it.Index(), v)
	return true
}

func (it *bit) Get() uint {
	return it.buf.get0(it.Index())
}

func (it *bit) Index() int {
	return it.idx
}

func (it *bit) String() string {
	return int2String(it.Get(), B10)
}
