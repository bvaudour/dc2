package style

type Color string

const (
	NOCOL    Color = ""
	BLACK    Color = "232"
	RED      Color = "1"
	GREEN    Color = "2"
	YELLOW   Color = "3"
	BLUE     Color = "4"
	MAJENTA  Color = "5"
	CYAN     Color = "6"
	WHITE    Color = "7"
	LBLACK   Color = "8"
	LRED     Color = "9"
	LGREEN   Color = "10"
	LYELLOW  Color = "11"
	LBLUE    Color = "12"
	LMAJENTA Color = "13"
	LCYAN    Color = "14"
	LWHITE   Color = "15"
)

func (c Color) Foreground() string {
	if c == NOCOL {
		return ""
	}
	return "38;5;" + string(c)
}

func (c Color) Background() string {
	if c == NOCOL {
		return ""
	}
	return "48;5;" + string(c)
}
