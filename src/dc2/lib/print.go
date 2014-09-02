package lib

import (
	"fmt"
	stl "style"
	u "util"
)

func SPrint(s string) {
	stl.NewColor(stl.LGREEN, stl.BOLD).Print(s)
}

func MPrint(m *Message) {
	var c stl.Color
	switch m.tpe {
	case E:
		c = stl.LRED
	case W:
		c = stl.LYELLOW
	default:
		c = stl.LWHITE
	}
	stl.NewColor(c, stl.BOLD).Print(m.msg)
}

func HTitlePrint(s string) {
	stl.NewColor(stl.LRED, stl.BOLD).Print(s)
}

func HOptionPrint(o, h string) {
	o = "  " + o
	l := len(o)
	switch {
	case l > 30:
		o += "\\n"
		h = u.Repeat(" ", 31) + h
	case l < 30:
		h = u.Repeat(" ", 31-l) + h
	default:
		h = " " + h
	}
	fo := stl.NewStyle(stl.BOLD).Format(o)
	fh := stl.NewColor(stl.LBLUE, stl.BOLD).Format(h)
	fmt.Println(fo + fh)
}
