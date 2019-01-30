package levelone

import "fmt"

func FactorialDigitSum(n int) int {
	var prod = []int{1}
	for i := 2; i <= n; i++ {
		prod = product(i, prod)
	}
	s := 0
	fmt.Printf("Product : %v\n", prod)
	for _, e := range prod {
		s += e
	}
	return s
}

func product(n int, prod []int) []int {
	p := 0
	for j := len(prod) - 1; j >= 0; j-- {
		p = (prod[j] * n) + p
		prod[j] = p % 10
		p = p / 10
	}
	for p != 0 {
		prod = append([]int{p % 10}, prod...)
		p = p / 10
	}
	//fmt.Printf("prod: %v\n", prod)
	return prod
}
