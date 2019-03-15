package leveltwo

import "fmt"

func GetConsectivePrimeFactorTerm(cons int) int {
	count := 0
	i := 9
	s := 9
	pfMap := map[int]int{}
	p := 2
	for count != cons && p < i {
		n := i
		p = 2
		pfactors := 0
		for n != 1 && p < i {
			c := 0
			for n%p == 0 {
				c++
				n = n / p
			}
			if c != 0 {
				pfMap[p] = c
			}
			//fmt.Printf("n: %d p: %d count %d factors: %v\n", n, p, c, pfMap)
			pfactors = len(pfMap)
			p = findNextPrime(p)
		}
		//fmt.Printf("term %d has %v factors\n", i, pfMap)
		i++
		pfMap = map[int]int{}
		if pfactors == cons {
			count++
		} else {
			s = i
			count = 0
		}

	}
	fmt.Printf("Factors %d, %d, %d, %d\n", s, s+1, s+2, s+3)
	return s
}
