package levelone

import "fmt"

func FindCollatzSeqCount(limit int) int {
	max := -1
	for i := 13; i < limit; i++ {
		c := CollatzCount(i)
		if c > max {
			fmt.Printf("Max count is %d for value %d\n", c, i)
			max = c
		}
	}
	return max
}

func CollatzCount(n int) int {
	count := 1
	for n != 1 {
		count++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return count
}
