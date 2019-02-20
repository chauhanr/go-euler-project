package leveltwo

import (
	"fmt"
	"math"
)

type Triplet struct {
	a int
	b int
	c int
}

func initTr(tr []int) {
	for i := 0; i < len(tr); i++ {
		tr[i] = 0
	}
}

func FindMaxTrpiletCount(p int) int {
	tr := make([]int, p+1)
	initTr(tr)
	for a := 1; a < p; a++ {
		for b := a; b < p && b >= a; b++ {
			c := Sqrt(a, b)
			if (a*a)+(b*b) == c*c {
				peri := a + b + c
				if (a + b + c) <= p {
					fmt.Printf("solution (%d,%d,%d) peri - %d\n", a, b, c, peri)
					tr[peri]++
				} else {
					break
				}
			}
		}
	}
	maxTriplet := 0
	per := 0

	for i := 0; i < len(tr); i++ {
		if tr[i] > maxTriplet {
			maxTriplet = tr[i]
			per = i
		}
	}

	fmt.Printf("Max Triplet count is %d at perimeter %d\n", maxTriplet, per)
	return maxTriplet
}

func Sqrt(a, b int) int {
	n := a*a + b*b
	sq := math.Sqrt(float64(n))
	return int(sq)
}
