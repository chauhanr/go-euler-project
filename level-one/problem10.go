package levelone

import "fmt"

func SumPrimeBelow(limit int64) int64 {
	var sum int64
	sum = 5
	var lPrime int64
	lPrime = 2
	k := 1

	var p1, p2 int64
	p1 = int64(6*k - 1)

	for p1 <= limit {
		if isPrimeInt(p1) {
			sum += p1
			lPrime = p1
		}
		p2 = int64(6*k + 1)

		if p2 <= limit && isPrimeInt(p2) {
			sum += p2
			lPrime = p2
		}
		k++
		p1 = int64(6*k - 1)
	}
	fmt.Printf("last prime before 2m %d\n", lPrime)
	return sum
}
