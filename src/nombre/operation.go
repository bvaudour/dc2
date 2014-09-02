package nombre

func Compare(n1, n2 Number) Number {
	return Int2Entier(n1.Compare(n2))
}

func Eq(n1, n2 Number) Number {
	return Bool2Entier(n1.Eq(n2))
}

func Ne(n1, n2 Number) Number {
	return Bool2Entier(n1.Ne(n2))
}

func Ge(n1, n2 Number) Number {
	return Bool2Entier(n1.Ge(n2))
}

func Le(n1, n2 Number) Number {
	return Bool2Entier(n1.Le(n2))
}

func Gt(n1, n2 Number) Number {
	return Bool2Entier(n1.Gt(n2))
}

func Lt(n1, n2 Number) Number {
	return Bool2Entier(n1.Lt(n2))
}

func Plus(n1, n2 Number) Number {
	switch {
	case n1.IsDecimal() || n2.IsDecimal():
		return n1.Decimal().Plus(n2.Decimal())
	case n1.IsBase() || n2.IsBase():
		return n1.Base().Plus(n2.Base())
	default:
		return n1.Entier().Plus(n2.Entier())
	}
}

func CEntier(n Number) Number {
	return n.Entier()
}

func CDecimal(n Number) Number {
	return n.Decimal()
}

func CBase2(n Number) Number {
	return n.Base().SetBase(2)
}

func CBase10(n Number) Number {
	return n.Base().SetBase(10)
}

func CBase16(n Number) Number {
	return n.Base().SetBase(16)
}

func Abs(n Number) Number {
	switch {
	case n.IsDecimal():
		return n.Decimal().Abs()
	case n.IsBase():
		return n.Base().Abs()
	default:
		return n.Entier().Abs()
	}
}

func Int(n Number) int {
	e := n.Entier()
	if e.buf.Len() == 1 && e.Positive() {
		return int(e.buf.Get(-1))
	}
	return -1
}

func Minus(n1, n2 Number) Number {
	switch {
	case n1.IsDecimal() || n2.IsDecimal():
		return n1.Decimal().Minus(n2.Decimal())
	case n1.IsBase() || n2.IsBase():
		return n1.Base().Minus(n2.Base())
	default:
		return n1.Entier().Minus(n2.Entier())
	}
}

func Time(n1, n2 Number) Number {
	switch {
	case n1.IsDecimal() || n2.IsDecimal():
		return n1.Decimal().Time(n2.Decimal())
	case n1.IsBase() || n2.IsBase():
		return n1.Base().Time(n2.Base())
	default:
		return n1.Entier().Time(n2.Entier())
	}
}

func By(n1, n2 Number) Number {
	switch {
	case n1.IsDecimal() || n2.IsDecimal():
		return n1.Decimal().By(n2.Decimal())
	case n1.IsBase() || n2.IsBase():
		return n1.Base().By(n2.Base())
	default:
		out := n1.Decimal().By(n2.Decimal())
		if out.IsInt() {
			return out.Entier()
		}
		return out
	}
}

func Mod(n1, n2 Number) Number {
	switch {
	case n1.IsDecimal() || n2.IsDecimal():
		return n1.Decimal().Mod(n2.Decimal())
	case n1.IsBase() || n2.IsBase():
		return n1.Base().Mod(n2.Base())
	default:
		return n1.Entier().Mod(n2.Entier())
	}
}

func DivMod(n1, n2 Number) (Number, Number) {
	switch {
	case n1.IsDecimal() || n2.IsDecimal():
		o1, o2 := n1.Decimal().DivMod(n2.Decimal())
		return o1.Entier(), o2
	case n1.IsBase() || n2.IsBase():
		return n1.Base().DivMod(n2.Base())
	default:
		return n1.Entier().DivMod(n2.Entier())
	}
}

func Pow(n1, n2 Number) Number {
	switch {
	case n1.IsDecimal() || n2.IsDecimal():
		return n1.Decimal().Pow(n2.Decimal())
	case n1.IsBase() || n2.IsBase():
		return n1.Base().Pow(n2.Base())
	default:
		out := n1.Decimal().Pow(n2.Decimal())
		if out.IsInt() {
			return out.Entier()
		}
		return out
	}
}

func Sqrt(n Number) Number {
	switch {
	case n.IsDecimal():
		return n.Decimal().Sqrt()
	case n.IsBase():
		return n.Base().Sqrt()
	default:
		out := n.Decimal().Sqrt()
		if out.IsInt() {
			return out.Entier()
		}
		return out
	}
}

func Fact(n Number) Number {
	switch {
	case n.IsDecimal():
		if n.Decimal().IsInt() {
			return n.Entier().Fact()
		}
		return EInf()
	case n.IsBase():
		return n.Base().Fact()
	default:
		return n.Entier().Fact()
	}
}

func Format(n Number) Number {
	if n.IsDecimal() {
		return n.Decimal().format()
	}
	return n
}
