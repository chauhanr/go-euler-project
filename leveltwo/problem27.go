package leveltwo

import (
	"fmt"
	"math"
)

func CalculateQuadCoefProd() int {
	lr := 0
	la := 0
	lb := 0

	b := findNextPrime(1)
	for b < 1000 {
		for a := -999; a < 1000; a++ {
			r := LongestRunQuadValue(a, b)
			if r > lr {
				lr = r
				la = a
				lb = b
			}
		}
		b = findNextPrime(b)
	}
	fmt.Printf("Longest Run is for (%d, %d) and is %d\n", la, lb, lr)
	return la * lb

}

func LongestRunQuadValue(a, b int) int {
	run := 0
	for n := 0; ; n++ {
		r := (n * n) + a*n + b
		if r < 0 {
			break
		}
		if IsPrimeInt(int64(r)) {
			run++
		} else {
			break
		}
	}
	return run
}

/*
  this function will find the next prime after the number that is provided as parameter
  the method will return a -1 if the input number of < 0 indicating that the negative
  have a limit to the maximum number of numbers the algo with try and that is 1000
*/
func findNextPrime(n int) int {
	d := n / 3
	p1 := 3*d - 1
	p2 := 3*d + 1
	np := -1
	//fmt.Printf("initial d for n: %d is %d\n", n, d)

	for np == -1 {
		if p1 >= 0 {
			if IsPrimeInt(int64(p1)) && p1 > n {
				np = p1
				break
			}
		}

		if IsPrimeInt(int64(p2)) && p2 > n {
			np = p2
		}
		d++
		p1 = 3*d - 1
		p2 = 3*d + 1
	}
	//fmt.Printf("next prime for %d is %d\n", n, np)
	return np
}

func IsPrimeInt(a int64) bool {
	sq := math.Sqrt(float64(a))
	sqInt := int64(sq)

	for sqInt > 1 {
		if a%sqInt == 0 {
			return false
		}
		sqInt--
	}
	return true
}
