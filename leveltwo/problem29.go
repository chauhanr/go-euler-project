package leveltwo

import "fmt"

func primeBelow(n int) int {
	p := []int{2}
	for i := 3; i < n; i++ {
		if IsPrimeInt(int64(i)) {
			p = append(p, i)
		}
	}
	fmt.Printf("primes %v\n", p)
	return len(p)
}
