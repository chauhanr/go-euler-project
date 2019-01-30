package levelone

import "fmt"

var dStore = make([]int, 100000)

func initStore() {
	for i := 0; i < len(dStore); i++ {
		dStore[i] = 0
		if isPrime(uint64(i)) {
			dStore[i] = 1
		}
	}
	dStore[1] = 1
}

type APair struct {
	a int
	b int
}

func FindAmicablePairsSum(n int) int {
	initStore()
	var pairs = make([]APair, 1)
	for i := 2; i < n; i++ {
		if dStore[i] == 1 {
			// do nothing
		} else if dStore[i] == 0 {
			sum := getDivisorsSum(i)
			dStore[i] = sum
			if dStore[sum] > 1 {
				if dStore[sum] == i && dStore[i] == sum {
					if sum != i {
						pairs = append(pairs, APair{i, sum})
					}
				}
			} else if dStore[sum] == 0 {
				s := getDivisorsSum(sum)
				dStore[sum] = s
				//fmt.Printf("sum and i store values (%d, %d)\n", dStore[sum], dStore[i])
				if dStore[sum] == i && dStore[i] == sum {
					if i != sum {
						fmt.Printf("Amicable Pair (%d, %d)\n", i, sum)
						pairs = append(pairs, APair{i, sum})
					}
				}
			}
		}
	}
	fmt.Printf("Pairs %v \n", pairs)
	sum := 0
	for _, p := range pairs {
		sum = sum + p.a + p.b
	}
	return sum
}

func getDivisorsSum(n int) int {
	d := []int{1}
	if isPrimeInt(int64(n)) {
		return 1
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			d = append(d, i)
		}
	}
	s := 0
	for _, e := range d {
		s += e
	}
	//fmt.Printf("divisors for %d :  %v, sum %d \n", n, d, s)
	return s
}
