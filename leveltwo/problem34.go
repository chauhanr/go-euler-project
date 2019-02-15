package leveltwo

import "fmt"

var facts []int64
var MAX int64

func initFacts() {
	facts = []int64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MAX = 0

	for i := 1; i <= 9; i++ {
		facts[i] = facts[i-1] * int64(i)
		MAX += facts[i]
	}
	fmt.Printf("facts %v MAX: %d\n", facts, MAX)
}

func FindAllFactSumN() int64 {
	if len(facts) == 0 {
		initFacts()
	}
	var sum int64
	sum = 0
	var i int64
	for i = 10; i <= MAX; i++ {
		n := GetDigitFactorialSum(i)
		if n == i {
			fmt.Printf("number %d fits criteria\n", n)
			sum += n
		}
	}
	return sum
}

func GetDigitFactorialSum(n int64) int64 {
	if len(facts) == 0 {
		initFacts()
	}
	var sum int64
	sum = 0
	for n != 0 {
		d := n % 10
		sum += facts[d]
		n = n / 10
	}
	return sum
}
