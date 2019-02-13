package levelone

import "fmt"

type PrimeFactor struct {
	factor   int64
	count    int
	multiple int64
}

func FindSmallestMultiple(value int64) int64 {

	var primes []PrimeFactor
	primes = append(primes, PrimeFactor{2, 0, 1})
	primes = append(primes, PrimeFactor{3, 0, 1})
	var k int64

	for k = 1; 6*k-1 < value; k++ {
		if isPrimeInt(6*k - 1) {
			primes = append(primes, PrimeFactor{6*k - 1, 0, 1})
		}
		if isPrimeInt(6*k + 1) {
			primes = append(primes, PrimeFactor{6*k + 1, 0, 1})
		}
	}

	var smallestMul int64
	smallestMul = 1
	var j int64
	j = 2

	for ; j <= value; j++ {
		val := j
		for i, prime := range primes {
			for val%prime.factor == 0 && val >= prime.factor {
				if smallestMul%j != 0 {
					prime.count++
					prime.multiple = prime.multiple * prime.factor
					primes[i] = prime
					smallestMul = smallestMul * prime.factor
				}
				val = val / prime.factor
			}
		}
	}
	fmt.Printf("Smallest Multiple is %d\n", smallestMul)
	return smallestMul
}

func factorial(limit int64) int64 {
	var fact int64
	var i int64
	fact = 1
	for i = 2; i <= limit; i++ {
		fact = fact * i
	}
	return fact
}
