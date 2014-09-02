package lib

import (
	u "util"
)

type Stack []Any

func smake(l, c int) Stack {
	return make([]Any, l, c)
}

func sappend(s Stack, aa ...Any) Stack {
	return append(s, aa...)
}

func scopy(s Stack) Stack {
	out := make([]Any, len(s), cap(s))
	copy(out, s)
	return out
}

func SInit() Stack {
	return smake(0, 10)
}

func (s *Stack) Push(aa ...Any) {
	*s = sappend(*s, aa...)
}

func (s *Stack) Remove(i int) {
	*s = (*s)[:len(*s)-i]
}

func (s *Stack) Last() Any {
	return (*s)[len(*s)-1]
}

func (s *Stack) Last2() (Any, Any) {
	return (*s)[len(*s)-2], (*s)[len(*s)-1]
}

func (s *Stack) Pop() Any {
	out := s.Last()
	s.Remove(1)
	return out
}

func (s *Stack) Pop2() (Any, Any) {
	out1, out2 := s.Last2()
	s.Remove(2)
	return out1, out2
}

func (s *Stack) Clear() {
	*s = SInit()
}

func (s Stack) Copy() Stack {
	return scopy(s)
}

func (s *Stack) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Stack) Reverse() {
	l := len(*s)
	for i := 0; i < l/2; i++ {
		s.Swap(i, l-1-i)
	}
}

func (s *Stack) Reverse2() {
	l := len(*s)
	s.Swap(l-1, l-2)
}

func (s *Stack) Mkstring(f func(Any) string) string {
	f2 := func(x interface{}) string { return f(x.(Any)) }
	return u.SliceStringBy(*s, f2)
}

func (s *Stack) Str() string {
	return s.Mkstring(AStr)
}

func (s *Stack) Sci() string {
	return s.Mkstring(ASci)
}

func (s *Stack) Format() {
	for i, a := range *s {
		(*s)[i] = AFormat(a)
	}
}
