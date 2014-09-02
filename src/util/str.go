package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Concat(acc *string, sep, sstr string) string {
	switch {
	case *acc == "":
		*acc = sstr
	case sstr != "":
		*acc += (sep + sstr)
	}
	return *acc
}

func Repeat(s string, l int) string {
	return strings.Repeat(s, l)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func StringOf(i interface{}) string {
	return Format("%v", i)
}

func Format(form string, i ...interface{}) string {
	return fmt.Sprintf(form, i...)
}

func String2Array(s string) []string {
	return strings.Fields(s)
}

func Count(s, subs string) int {
	return strings.Count(s, subs)
}

func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

func Join(s []string, sep string) string {
	return strings.Join(s, sep)
}

func Range(s string) []int {
	out := make([]int, 0)
	r := regexp.MustCompile("\\d+(-\\d+)?(,\\d+(-\\d+)?)*")
	if !r.MatchString(s) {
		return out
	}
	vg := Split(s, ",")
	for _, v := range vg {
		rg := Split(v, "-")
		if len(rg) == 1 {
			i, _ := strconv.Atoi(rg[0])
			out = append(out, i)
		} else {
			i1, _ := strconv.Atoi(rg[0])
			i2, _ := strconv.Atoi(rg[1])
			if i1 >= i2 {
				for i := i1; i >= i2; i-- {
					out = append(out, i)
				}
			} else {
				for i := i1; i <= i2; i++ {
					out = append(out, i)
				}
			}
		}
	}
	return out
}
