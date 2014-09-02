package lib

import (
	"fmt"
	u "util"
)

var DOp = map[string]func(*Dc, string) *Message{
	"v":   DC1to1,
	"!":   DC1to1,
	"|":   DC1to1,
	"cI":  DC1to1,
	"cD":  DC1to1,
	"c2":  DC1to1,
	"c10": DC1to1,
	"c16": DC1to1,
	"+":   DC2to1,
	"-":   DC2to1,
	"*":   DC2to1,
	"/":   DC2to1,
	"%":   DC2to1,
	"^":   DC2to1,
	"<=>": DC2to1,
	"=":   DC2to1,
	"<>":  DC2to1,
	">=":  DC2to1,
	"<=":  DC2to1,
	">":   DC2to1,
	"<":   DC2to1,
	"~":   DC2to2,
}

var D0 = map[string]func(*Dc) *Message{
	"h":  HDisplay,
	"hr": HRDisplay,
	"o":  ODisplay,
	"k":  OAuto,
	"K":  OFixe,
	"a":  SLConv,
	"p":  SLDisplay,
	"e":  SLSci,
	"c":  SLClr,
	"d":  SLDup,
	"r":  SLRev,
	"P":  SDisplay,
	"E":  SSci,
	"C":  SClr,
	"D":  SDup,
	"R":  SRev,
	":P": MDisplay,
	":E": MSci,
}

var DKey = map[string]func(*Dc, string) *Message{
	":c": RLClr,
	":d": RLDup,
	":r": RLRev,
	":P": RDisplay,
	":E": RSci,
	":C": RClr,
	":D": RDup,
	":R": RRev,
	":l": RLUnstore,
	":s": RLStore,
	":L": RUnstore,
	":S": RStore,
}

var DKeyOp = map[string]func(*Dc, string, string) *Message{
	":=":  RCond,
	":<>": RCond,
	":>=": RCond,
	":<=": RCond,
	":>":  RCond,
	":<":  RCond,
}

func DC1to1(dc *Dc, s string) *Message {
	if len(*dc.stk) < 1 {
		return Error(EMPTYSTACK)
	}
	r, m := C1to1(dc.stk.Pop(), s)
	if m != nil {
		return m
	}
	dc.stk.Push(r)
	return nil
}

func DC2to1(dc *Dc, s string) *Message {
	if len(*dc.stk) < 2 {
		return Error(u.Format(EMPTYSTACKN, 2))
	}
	a1, a2 := dc.stk.Pop2()
	r, m := C2to1(a1, a2, s)
	if m != nil {
		return m
	}
	dc.stk.Push(r)
	return nil
}

func DC2to2(dc *Dc, s string) *Message {
	if len(*dc.stk) < 2 {
		return Error(u.Format(EMPTYSTACKN, 2))
	}
	a1, a2 := dc.stk.Pop2()
	r1, r2, m := C2to2(a1, a2, s)
	if m != nil {
		return m
	}
	dc.stk.Push(r1, r2)
	return nil
}

func HDisplay(dc *Dc) *Message {
	PrintAll()
	return nil
}

func HRDisplay(dc *Dc) *Message {
	PrintRubriques()
	return nil
}

func HSDisplay(dc *Dc, s string) *Message {
	if PrintSections(u.Range(s)) {
		return nil
	}
	return Warning(NOSECTION)
}

func OAuto(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	if m := OSetAuto(dc.stk.Pop()); m != nil {
		return m
	}
	dc.stk.Format()
	for _, s := range *dc.mem {
		s.Format()
	}
	return nil
}

func OFixe(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	if m := OSetFixe(dc.stk.Pop()); m != nil {
		return m
	}
	dc.stk.Format()
	for _, s := range *dc.mem {
		s.Format()
	}
	return nil
}

func ODisplay(dc *Dc) *Message {
	fmt.Println(OStr())
	return nil
}

func SLConv(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	if a := dc.stk.Pop(); IsNumber(a) {
		dc.stk.Push(AStr(a))
	} else {
		dc.stk.Push(a)
	}
	return nil
}

func SLDisplay(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	SPrint(AStr(dc.stk.Last()))
	return nil
}

func SDisplay(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	SPrint(dc.stk.Str())
	return nil
}

func SLSci(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	SPrint(ASci(dc.stk.Last()))
	return nil
}

func SSci(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	SPrint(dc.stk.Sci())
	return nil
}

func MDisplay(dc *Dc) *Message {
	if len(*dc.mem) < 1 {
		return Warning(EMPTYMEMORY)
	}
	SPrint(dc.mem.Str0())
	return nil
}

func MSci(dc *Dc) *Message {
	if len(*dc.mem) < 1 {
		return Warning(EMPTYMEMORY)
	}
	SPrint(dc.mem.Sci0())
	return nil
}

func SLClr(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	dc.stk.Remove(1)
	return nil
}

func SClr(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	dc.stk.Clear()
	return nil
}

func SLDup(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	dc.stk.Push(dc.stk.Last())
	return nil
}

func SDup(dc *Dc) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	dc.stk.Push(dc.stk.Copy()...)
	return nil
}

func SLRev(dc *Dc) *Message {
	if len(*dc.stk) < 2 {
		return Warning(u.Format(EMPTYSTACKN, 2))
	}
	dc.stk.Reverse2()
	return nil
}

func SRev(dc *Dc) *Message {
	if len(*dc.stk) < 2 {
		return Warning(u.Format(EMPTYSTACKN, 2))
	}
	dc.stk.Reverse()
	return nil
}

func RDisplay(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	SPrint(dc.mem.Str(k))
	return nil
}

func RSci(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	SPrint(dc.mem.Sci(k))
	return nil
}

func RLClr(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	s := dc.mem.Get(k)
	if len(*s) == 1 {
		dc.mem.Remove(k)
	} else {
		s.Remove(1)
	}
	return nil
}

func RClr(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	dc.mem.Remove(k)
	return nil
}

func RLDup(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	dc.mem.Add(k, dc.mem.Last(k))
	return nil
}

func RDup(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	dc.mem.Add(k, dc.mem.Copy(k))
	return nil
}

func RLRev(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	s := dc.mem.Get(k)
	if len(*s) < 2 {
		return Warning(u.Format(EMPTYREGISTERN, k, 2))
	}
	s.Reverse2()
	return nil
}

func RRev(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	s := dc.mem.Get(k)
	if len(*s) < 2 {
		return Warning(u.Format(EMPTYREGISTERN, k, 2))
	}
	s.Reverse()
	return nil
}

func RLUnstore(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	dc.stk.Push(dc.mem.Last(k))
	return nil
}

func RUnstore(dc *Dc, k string) *Message {
	if !dc.mem.Exists(k) {
		return Warning(u.Format(EMPTYREGISTER, k))
	}
	dc.stk.Push(dc.mem.Copy(k)...)
	return nil
}

func RLStore(dc *Dc, k string) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	e := dc.stk.Last()
	if dc.mem.Exists(k) {
		MPrint(Info(u.Format(NONEMPTYREGISTER, k)))
		switch u.Question3(ERASEREGISTER, u.YES) {
		case u.YES:
			dc.mem.Put(k, e)
		case u.NO:
			dc.mem.Add(k, e)
		default:
			return Info(CANCEL)
		}
	} else {
		dc.mem.Put(k, e)
	}
	dc.stk.Remove(1)
	return Info(u.Format(CREATEDREGISTER, k))
}

func RStore(dc *Dc, k string) *Message {
	if len(*dc.stk) < 1 {
		return Warning(EMPTYSTACK)
	}
	e := dc.stk.Copy()
	if dc.mem.Exists(k) {
		MPrint(Info(u.Format(NONEMPTYREGISTER, k)))
		switch u.Question3(ERASEREGISTER, u.YES) {
		case u.YES:
			dc.mem.Put(k, e...)
		case u.NO:
			dc.mem.Add(k, e...)
		default:
			return Info(CANCEL)
		}
	} else {
		dc.mem.Put(k, e...)
	}
	dc.stk.Clear()
	return Info(u.Format(CREATEDREGISTER, k))
}

func RCond(dc *Dc, k, s string) *Message {
	if e := DC2to1(dc, s); e != nil {
		return e
	}
	a := N(dc.stk.Pop())
	if a.Is(1) {
		return RUnstore(dc, k)
	}
	return nil
}

func DNombre(dc *Dc, s string) *Message {
	a := String2Number(s)
	if a.None() {
		return Error(u.Format(UNKNOWNCOMMAND, s))
	}
	dc.stk.Push(a)
	return nil
}
