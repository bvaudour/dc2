package main

import (
	dc2 "dc2/lib"
	u "util"
)

func main() {
	dc := dc2.New()
	for ok := true; ok; {
		s := u.Split(u.Scanln("> "), " ")
		ok = dc.Parse(s)
	}
}
