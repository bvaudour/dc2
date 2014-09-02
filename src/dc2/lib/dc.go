package lib

import (
	u "util"
)

type Dc struct {
	stk *Stack
	mem *Memory
}

func New() *Dc {
	s, m := SInit(), MInit()
	return &Dc{&s, &m}
}

func (dc *Dc) Parse(args []string) bool {
	isChaine, nbBr, chaine := false, 0, ""
	for _, s := range args {
		switch {
		case s == "":
			continue
		case !isChaine && s[0] == '[':
			isChaine = true
			fallthrough
		case isChaine:
			nbBr += (u.Count(s, "[") - u.Count(s, "]"))
			u.Concat(&chaine, " ", s)
			if nbBr < 0 {
				MPrint(Error(u.Format(UNKNOWNCOMMAND, chaine)))
				dc.stk.Clear()
				isChaine = false
				nbBr = 0
				chaine = ""
			} else if nbBr == 0 {
				if l := len(chaine); l > 2 {
					dc.stk.Push(chaine[1 : l-1])
				}
				isChaine = false
				chaine = ""
			}
		case s[0] == '#':
			break
		case s == "q":
			return false
		case s == "x":
			if m, end := dc.execLst(); end {
				return end
			} else if m != nil {
				MPrint(m)
				if m.IsError() {
					dc.stk.Clear()
				}
			}
		default:
			if m := dc.parseElt(s); m != nil {
				MPrint(m)
				if m.IsError() {
					dc.stk.Clear()
				}
			}
		}
	}
	if isChaine {
		MPrint(Error(u.Format(UNKNOWNCOMMAND, chaine)))
		dc.stk.Clear()
	}
	return true
}

func (dc *Dc) parseElt(s string) *Message {
	switch {
	case u.Contains(D0, s):
		return D0[s](dc)
	case u.Contains(DOp, s):
		return DOp[s](dc, s)
	case s[0] == 'h':
		return HSDisplay(dc, s[1:])
	case len(s) > 2 && u.Contains(DKey, s[:2]):
		return DKey[s[:2]](dc, s[2:])
	case len(s) > 2 && u.Contains(DKeyOp, s[:2]):
		return DKeyOp[s[:2]](dc, s[2:], s[1:2])
	case len(s) > 3 && u.Contains(DKeyOp, s[:3]):
		return DKeyOp[s[:3]](dc, s[3:], s[1:3])
	default:
		return DNombre(dc, s)
	}
}

func (dc *Dc) execLst() (m *Message, end bool) {
	if len(*dc.stk) == 0 {
		m = Warning(EMPTYSTACK)
	} else {
		a := dc.stk.Pop()
		if IsNumber(a) {
			dc.stk.Push(a)
		} else {
			end = dc.Parse(u.String2Array(S(a)))
		}
	}
	return
}
