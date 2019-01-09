package levelone

import (
	"fmt"
	"math"
)

func LargestPrimeFactor2(value int64) int64 {
	var maxPrime int64 = -1
	for value%2 == 0 {
		maxPrime = 2
		value = value / 2
	}
	var i int64
	for i = 3; i <= Sqrt(value); i += 2 {
		for value%i == 0 {
			maxPrime = i
			value = value / i
		}
		//fmt.Printf("Value :%d |  i: %d |  maxPrime: %d\n", value, i, maxPrime)
	}

	if maxPrime == -1 || value != 1 {
		maxPrime = value
	}

	return maxPrime
}

func Sqrt(a int64) int64 {
	sq := math.Sqrt(float64(a))
	sqInt := int64(sq)
	return sqInt
}

func LargestPrimeFactor(value uint64) uint64 {
	var pf, k uint64
	pf = 0
	k = determineLargestK(value)

	fmt.Printf("K for value %d is %d \n", value, k)

	if value%2 == 0 {
		pf = 2
	}
	if value%3 == 0 {
		pf = 3
	}

	n1, n2 := next2Elements(k)
	for k != 0 {
		if value%n2 == 0 {
			if isPrime(n2) {
				pf = n2
				return pf
			}
		}
		if value%n1 == 0 {
			if isPrime(n1) {
				pf = n1
				return pf
			}
		}
		k--
		n1, n2 = next2Elements(k)
		//fmt.Printf("n1 : %d, n2 %d\n", n1, n2)
	}
	if pf == 0 {
		pf = value
	}
	return pf

}

func next2Elements(k uint64) (uint64, uint64) {
	n1 := 6*k - 1
	n2 := 6*k + 1
	return n1, n2
}

func determineLargestK(value uint64) uint64 {
	var k, h uint64
	k = value / 6
	h = value / 2
	for 6*k+1 > h && 6*k-1 > h {
		k--
	}
	return k
}

func isPrime(a uint64) bool {
	sq := math.Sqrt(float64(a))
	sqInt := uint64(sq)

	for sqInt > 1 {
		if a%sqInt == 0 {
			return false
		}
		sqInt--
	}
	return true
}
