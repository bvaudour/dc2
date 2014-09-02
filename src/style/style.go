package style

import (
	"fmt"
	u "util"
)

type Style struct {
	fg Color
	bg Color
	ef []Effect
}

func NewColor(fg Color, ef ...Effect) *Style {
	return &Style{fg, NOCOL, u.Unic(ef).([]Effect)}
}

func NewStyle(ef ...Effect) *Style {
	return NewColor(NOCOL, ef...)
}

func End() *Style {
	return NewStyle()
}

func (s *Style) SetBackground(bg Color) *Style {
	s.bg = bg
	return s
}

func (s *Style) SetForeground(fg Color) *Style {
	s.fg = fg
	return s
}

func (s *Style) AddEffect(e ...Effect) *Style {
	s.ef = u.Unic(append(s.ef, e...)).([]Effect)
	return s
}

func (s *Style) RemoveEffect(e ...Effect) *Style {
	ef := make([]interface{}, len(e))
	for i, v := range e {
		ef[i] = v
	}
	s.ef = u.Filter(s.ef, ef...).([]Effect)
	return s
}

func (s *Style) String() string {
	out := ""
	for _, e := range s.ef {
		u.Concat(&out, ";", e.String())
	}
	u.Concat(&out, ";", s.fg.Foreground())
	u.Concat(&out, ";", s.bg.Background())
	return "\033[" + out + "m"
}

func (s *Style) Format(e string) string {
	return s.String() + e + End().String()
}

func (s *Style) Print(e string) {
	fmt.Println(s.Format(e))
}

func print(st []*Style, str []string, sep string) {
	if len(st) != len(str) {
		panic("Style and string arrays must contain same number of element!")
	}
	e := ""
	for i, s := range st {
		u.Concat(&e, sep, s.Format(str[i]))
	}
	fmt.Println(e)
}

func Print(st []*Style, str []string) {
	print(st, str, "")
}

func PrintWithSpace(st []*Style, str []string) {
	print(st, str, " ")
}
