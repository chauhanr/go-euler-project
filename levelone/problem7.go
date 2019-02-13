package levelone

import "fmt"

/*
   function will return the nth prime based on the input sent to the method.
   so if you need the 101st prime just send 101
*/
func FindNthPrime(n int) int64 {
	if n == 1 {
		return 2
	} else if n == 2 {
		return 3
	}

	pCount := 2
	var k int64
	k = 1
	var nPrime int64
	nPrime = 2

	for pCount <= n {
		if isPrimeInt(6*k - 1) {
			pCount++
			nPrime = 6*k - 1
		}
		if pCount == n {
			break
		}
		if isPrimeInt(6*k + 1) {
			pCount++
			nPrime = 6*k + 1
		}
		k++
	}

	fmt.Printf("The %dth Prime is %d\n", pCount, nPrime)
	return nPrime

}
