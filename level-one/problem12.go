package levelone

import "fmt"

func CalculateMultipleCount() (int64, int64) {
	var n int64
	n = 2
	sum := SumN(n)
	count, sum := FactorCombCount(sum)
	for count < 500 {
		n++
		sum = SumN(n)
		count, sum = FactorCombCount(sum)
	}
	fmt.Printf("Triangle Number with 500 divisors is %d and count is %d\n", sum, count)
	return count, sum
}

func SumN(n int64) int64 {
	return int64(n * (n + 1) / 2)
}

func FactorCombCount(p int64) (mul int64, sum int64) {
	init := p
	var k int64
	k = 1
	var pFactors []PrimeFactor
	var f PrimeFactor
	f = PrimeFactor{factor: 2, count: 0}
	for p > 1 && p%2 == 0 {
		//fmt.Printf("Factor 2 found for %d\n", p)
		f.count++
		p = p / 2
	}
	if f.count != 0 {
		pFactors = append(pFactors, f)
	}
	f = PrimeFactor{factor: 3, count: 0}
	for p > 1 && p%3 == 0 {
		f.count++
		p = p / 3
	}
	if f.count != 0 {
		pFactors = append(pFactors, f)
	}

	for p > 1 && 6*k-1 < init {
		prime := 6*k - 1
		if isPrimeInt(prime) {
			f = PrimeFactor{factor: prime, count: 0}
			for p > 1 && p%prime == 0 {
				f.count++
				p = p / prime
			}
			if f.count != 0 {
				pFactors = append(pFactors, f)
			}
		}
		prime = 6*k + 1
		if prime < init && isPrimeInt(prime) {
			f = PrimeFactor{factor: prime, count: 0}
			for p > 1 && p%prime == 0 {
				f.count++
				p = p / prime
			}
			if f.count != 0 {
				pFactors = append(pFactors, f)
			}
		}
		k++
	}
	//fmt.Printf("Factors for %d is %v+\n", init, pFactors)
	// we have all the factors we need to calculate number of total multiples.
	mul = 1
	for _, fact := range pFactors {
		mul *= int64(fact.count + 1)
	}
	//fmt.Printf("Num of multiples for sum %d is %d\n", init, mul)
	return mul, init

}
