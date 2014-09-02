package lib

import (
	n "nombre"
	u "util"
)

var op1to1 = map[string]func(n.Number) n.Number{
	"v":   n.Sqrt,
	"!":   n.Fact,
	"|":   n.Abs,
	"cI":  n.CEntier,
	"cD":  n.CDecimal,
	"c2":  n.CBase2,
	"c10": n.CBase10,
	"c16": n.CBase16,
}

var op2to1 = map[string]func(n.Number, n.Number) n.Number{
	"+":   n.Plus,
	"-":   n.Minus,
	"*":   n.Time,
	"/":   n.By,
	"%":   n.Mod,
	"^":   n.Pow,
	"<=>": n.Compare,
	"=":   n.Eq,
	"<>":  n.Ne,
	">=":  n.Ge,
	"<=":  n.Le,
	">":   n.Gt,
	"<":   n.Lt,
}

var op2to2 = map[string]func(n.Number, n.Number) (n.Number, n.Number){
	"~": n.DivMod,
}

type Any interface{}

func String2Number(e string) n.Number {
	return n.Parse(e)
}

func Number2String(e n.Number) string {
	return e.String()
}

func IsNumber(a Any) bool {
	switch a.(type) {
	case n.Number:
		return true
	default:
		return false
	}
}

func IsString(a Any) bool {
	switch a.(type) {
	case string:
		return true
	default:
		return false
	}
}

func N(a Any) n.Number {
	return a.(n.Number)
}

func S(a Any) string {
	return a.(string)
}

func Number(a Any) (n.Number, bool) {
	if IsNumber(a) {
		return N(a), true
	}
	return nil, false
}

func String(a Any) (string, bool) {
	if IsNumber(a) {
		return S(a), true
	}
	return "", false
}

func AStr(a Any) string {
	if IsNumber(a) {
		return N(a).String()
	}
	return "'" + S(a) + "'"
}

func ASci(a Any) string {
	if IsNumber(a) {
		return N(a).Sci()
	}
	return "'" + S(a) + "'"
}

func C1to1(a Any, k string) (r n.Number, m *Message) {
	if !IsNumber(a) {
		m = Error(u.Format(NAN, AStr(a)))
	} else {
		r = op1to1[k](N(a))
		if r.None() {
			switch k {
			case "v":
				m = Error(NEGATIVESQRT)
			case "!":
				if r.Negative() {
					m = Error(NEGATIVEFACT)
				} else {
					m = Error(DECIMALFACT)
				}
			default:
				m = Error(UNKNOWN)
			}
		}
	}
	return
}

func C2to1(a1, a2 Any, k string) (r n.Number, m *Message) {
	if !IsNumber(a1) {
		m = Error(u.Format(NAN, AStr(a1)))
	} else if !IsNumber(a2) {
		m = Error(u.Format(NAN, AStr(a2)))
	} else {
		r = op2to1[k](N(a1), N(a2))
		if r.None() {
			switch k {
			case "^":
				fallthrough
			case "%":
				fallthrough
			case "/":
				m = Error(DIVIDEBYZERO)
			default:
				m = Error(UNKNOWN)
			}
		}
	}
	return
}

func C2to2(a1, a2 Any, k string) (r1, r2 n.Number, m *Message) {
	if !IsNumber(a1) {
		m = Error(u.Format(NAN, AStr(a1)))
	} else if !IsNumber(a2) {
		m = Error(u.Format(NAN, AStr(a2)))
	} else {
		r1, r2 = op2to2[k](N(a1), N(a2))
		if r1.None() || r2.None() {
			switch k {
			case "~":
				m = Error(DIVIDEBYZERO)
			default:
				m = Error(UNKNOWN)
			}
		}
	}
	return
}

func OSetAuto(a Any) *Message {
	if !IsNumber(a) {
		return Warning(u.Format(NAN, AStr(a)))
	}
	n.SetOption(n.Int(N(a)), false)
	return nil
}

func OSetFixe(a Any) *Message {
	if !IsNumber(a) {
		return Warning(u.Format(NAN, AStr(a)))
	}
	n.SetOption(n.Int(N(a)), true)
	return nil
}

func OStr() string {
	d, f := n.GetOption()
	fx := "Non"
	if f {
		fx = "Oui"
	}
	return u.Format("Précision: %d; Fixe: %s", d, fx)
}

func AFormat(a Any) Any {
	if IsNumber(a) {
		return n.Format(N(a))
	}
	return a
}
