package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	YES = iota
	NO
	CANCEL
)

func Scanln(invite string) string {
	fmt.Print(invite)
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return s[:len(s)-1]
}

func Question2(s string, def bool) bool {
	r := Lower(Scanln(s + " "))
	b := r != ""
	switch {
	case b && (r[0] == 'o' || r[0] == 'y'):
		return true
	case b && r[0] == 'n':
		return false
	default:
		return def
	}
}

func Question3(s string, def uint64) uint64 {
	r := Lower(Scanln(s + " "))
	b := r != ""
	switch {
	case b && (r[0] == 'o' || r[0] == 'y'):
		return YES
	case b && r[0] == 'n':
		return NO
	case b && (r[0] == 'a' || r[0] == 'c'):
		return CANCEL
	default:
		return def
	}
}

func ReadFile(path string) (lines []string, err error) {
	b, e := ioutil.ReadFile(path)
	if e != nil {
		err = e
		lines = []string{}
	} else {
		lines = Split(string(b), "\n")
	}
	return
}

func WriteFile(path string, lines []string) error {
	b := append([]byte(Join(lines, "\n")), '\n')
	return ioutil.WriteFile(path, b, 0644)
}

func AppendFile(path string, lines []string) error {
	r, _ := ReadFile(path)
	return WriteFile(path, append(r, lines...))
}
