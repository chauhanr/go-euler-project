package leveltwo

import (
	"fmt"
)

func FindLongestDecimalCount(l int) (max int, n int) {
	max = 0
	for i := 2; i < l; i++ {
		rec := FindRecurringDecimalCount(i)
		if rec > max {
			max = rec
			n = i
		}
	}
	fmt.Printf("max: %d, n: %d\n", max, n)
	return max, n
}

func FindRecurringDecimalCount(dem int) int {
	a := []int{}
	rem := map[int]int{}
	num := 1
	for num != 0 {
		for num < dem {
			num *= 10
			if num < dem {
				if _, ok := rem[num]; ok {
					break
				} else {
					a = append(a, 0)
					rem[num] = 1
				}
			}
		}
		q := num / dem
		num = num % dem
		if _, ok := rem[num]; ok {
			break
		} else {
			rem[num] = 1
		}
		a = append(a, q)
	}
	//fmt.Printf("array %v\n", a)
	//fmt.Printf("Recurring decimal count for number %d is %d\n", dem, len(rem))

	if num == 0 {
		return 0
	}
	return len(rem)
}
