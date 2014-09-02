package style

type Effect string

const (
	NORMAL    Effect = "0"
	BOLD      Effect = "1"
	ITALIC    Effect = "3"
	UNDERLINE Effect = "4"
	BLINK     Effect = "5"
	INVERTED  Effect = "7"
)

func (e Effect) String() string {
	return string(e)
}
