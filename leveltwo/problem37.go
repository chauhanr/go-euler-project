package leveltwo

import (
	"fmt"
	"math"
)

func GetNextPrime(n int64) int64 {
	m := n / 3
	for m > 0 {
		if (3*m - 1) > n {
			if IsPrimeInt(3*m - 1) {
				return 3*m - 1
			}
		}
		if 3*m+1 > n && IsPrimeInt(3*m+1) {
			return 3*m + 1
		}
		m++
	}
	return m
}

func isRightPrime(n int64) bool {
	d := n
	l := 0
	for d != 0 {
		d = d / 10
		l++
	}
	m := int64(math.Pow10(l))
	for m != 1 {
		p := n % m
		if !IsPrimeInt(p) {
			return false
		}
		m = m / 10
	}
	return true
}

func isLeftPrime(n int64) bool {
	d := n
	for d != 0 {
		if !IsPrimeInt(d) {
			return false
		}
		d = d / 10
	}
	return true
}

func FindLeftRightPrimeSum() int64 {
	var c, i int64
	c = 0
	i = GetNextPrime(10)
	var sum int64
	sum = 0
	for c != 11 {
		if isRightPrime(i) && isLeftPrime(i) {
			fmt.Printf("LR Prime %d\n", i)
			sum += i
			c++
		}
		i = GetNextPrime(i)
	}
	return sum
}
