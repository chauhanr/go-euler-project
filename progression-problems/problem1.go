package progression

import (
	"errors"
	"fmt"
)

func determineNthTerm(first int, diff int, nValue int) int {
	n := ((nValue - first) / diff) + 1
	return n
}

func determineLastTerm(finalValue int, divisor int) (int, error) {
	i := finalValue - 1
	for i%divisor != 0 && i >= 1 {
		i--
	}
	if i != 1 {
		return i, nil
	} else {
		return -1, errors.New("Cannot determine the last divisible value")
	}
}

func determineSum(v1, v2, last int) int {
	vlast1, err := determineLastTerm(last, v1)
	if err != nil {
		fmt.Errorf("Error determining the last term for %d on %d", v1, last)
	}
	vlast2, err := determineLastTerm(last, v2)
	if err != nil {
		fmt.Errorf("Error determining the last term for %d on %d", v2, last)
	}
	vlast3, err := determineLastTerm(last, v1*v2)
	if err != nil {
		fmt.Errorf("Error determining the last term for %d on %d", v1*v2, last)
	}

	nTerm1 := determineNthTerm(v1, v1, vlast1)
	nTerm2 := determineNthTerm(v2, v2, vlast2)
	nTerm3 := determineNthTerm(v1*v2, v1*v2, vlast3)
	v3 := v1 * v2

	sum1 := progressionSum(v1, v1, nTerm1)
	sum2 := progressionSum(v2, v2, nTerm2)
	sum3 := progressionSum(v3, v3, nTerm3)

	totalSum := sum1 + sum2 - sum3
	return totalSum
}

func progressionSum(a, d, n int) int {
	sum := (2*a*n + (n-1)*n*d) / 2
	return sum
}
