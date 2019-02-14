package leveltwo

import "fmt"

type Fraction struct {
	num int
	den int
}

func FindProductDen() int {
	r := []Fraction{}
	for i := 11; i < 100; i++ {
		for j := 99; j > i; j-- {
			f := Fraction{i, j}
			if isNonTrivialNum(f) {
				n := removeCommonDigit(f)
				if AreFractionsEqual(f, n) {
					fmt.Printf("Original %v, Reduced %v\n", f, n)
					r = append(r, n)
				}
			}
		}
	}

	fmt.Printf("Factions Collection %v\n", r)
	pd := 1
	pn := 1
	for _, v := range r {
		pd *= v.den
		pn *= v.num
	}
	for i := 2; i < 10; i++ {
		for pd%i == 0 && pn%i == 0 {
			pd = pd / i
			pn = pn / i
		}
	}

	return pd
}

func AreFractionsEqual(a, b Fraction) bool {
	if a.den == 0 || b.den == 0 {
		return false
	}
	var af float64
	var bf float64

	af = float64(a.num) / float64(a.den)
	bf = float64(b.num) / float64(b.den)

	if af == bf {
		return true
	} else {
		return false
	}
}

func isNonTrivialNum(a Fraction) bool {
	if a.num%11 == 0 && a.den%11 == 0 {
		return false
	}
	if a.num%10 == (a.den/10)%10 && (a.num/10)%10 == a.den%10 {
		return false
	}
	if a.num%10 == (a.den/10)%10 || (a.num/10)%10 == a.den%10 {
		return true
	}

	return false
}

func removeCommonDigit(a Fraction) Fraction {
	n := a.num
	for n != 0 {
		nc := n % 10
		n = n / 10
		d := a.den
		for d != 0 {
			dc := d % 10
			if dc == nc {
				f := modifyFraction(a, nc)
				return f
			}
			d = d / 10
		}
	}
	return Fraction{a.num, a.den}
}

func modifyFraction(a Fraction, n int) Fraction {
	f := Fraction{}
	if a.num%10 != n && (a.num/10)%10 != n {
		return Fraction{a.num, a.den}
	}

	if a.den%10 != n && (a.den/10)%10 != n {
		return Fraction{a.num, a.den}
	}

	if a.num%10 != n {
		f.num = a.num % 10
	} else {
		f.num = (a.num / 10) % 10
	}

	if a.den%10 != n {
		f.den = a.den % 10
	} else {
		f.den = (a.den / 10) % 10
	}
	return f
}
