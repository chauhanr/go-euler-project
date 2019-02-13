package levelone

import "fmt"

var abundants = []int{12}

func isNumberAbundant(n int) (bool, int) {
	s := getDivisorsSum(n)
	if s > n {
		return true, s
	}
	return false, -1
}

func findAllAbundants(limit int) {
	for i := 13; i < limit+1; i++ {
		a, _ := isNumberAbundant(i)
		if a {
			abundants = append(abundants, i)
		}
	}
	fmt.Printf("Abundants list count is  %d \n", len(abundants))
}

func FindAbundantSums() int {
	sum := 0
	limit := 28123
	findAllAbundants(limit)

	var canWrittenAsAbs = make([]bool, limit+1)
	for j := 0; j < len(canWrittenAsAbs); j++ {
		canWrittenAsAbs[j] = false
	}
	for x := 0; x < len(abundants); x++ {
		for y := x; y < len(abundants); y++ {
			if abundants[x]+abundants[y] <= limit {
				canWrittenAsAbs[abundants[x]+abundants[y]] = true
			}
		}
	}

	for i := 0; i < limit+1; i++ {
		if !canWrittenAsAbs[i] {
			sum += i
		}
	}

	//fmt.Printf("All abundant sums : %v\n", abundants)
	return sum
}
