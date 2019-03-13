package leveltwo

import (
	"fmt"
	"math"
)

func GetConvergedTerm() int {
	t := 286
	p := 165
	h := 143
	T := 40755
	P := 1
	H := 0

	for T != P && P != H {
		T = getTriangle(t)
		for T > P {
			p++
			P = getPentagonal(p)
		}
		if T == P {
			//fmt.Printf("For T and P are %d  equal for t %d p %d\n", T, t, p)
			H = getHectagonal(h)
			for P > H {
				h++
				H = getHectagonal(h)
			}
		}
		//fmt.Printf("(T, P, H) - (%d,%d,%d)\n", T, P, H)
		if T == P && H == P {
			fmt.Printf("(T,P,H) - (%d, %d, %d)\n", T, P, H)
			return T
		}
		p = 165
		h = 143
		P = 1
		H = 0
		t++
	}

	if T != P || T != H {
		return -1
	} else {
		return T
	}
}

func getTriangle(n int) int {
	r := n * (n + 1) / 2
	return r
}

func getPentagonal(n int) int {
	r := n * (3*n - 1) / 2
	return r
}

func getHectagonal(n int) int {
	r := n * (2*n - 1)
	return r
}

func isPentagonal(n int) bool {
	tr := math.Sqrt((1.0 + 24.0*float64(n))) + 1.0
	i := int64(tr)
	if i%6 == 0 {
		return true
	} else {
		return false
	}
}

func isTriangle(n int) bool {
	tr := 1.0 + math.Sqrt((1.0 + 8.0*float64(n)))
	i := int64(tr)
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}

func isHectagonal(n int) bool {
	tr := 1.0 + math.Sqrt((1.0 + 8.0*float64(n)))
	i := int64(tr)
	if tr == float64(i) {
		return true
	} else {
		return false
	}
}
