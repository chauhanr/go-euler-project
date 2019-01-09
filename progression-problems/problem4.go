package levelone

import (
	"fmt"
	"strconv"
)

func isPalindrome(value int) bool {
	str := strconv.Itoa(value)
	l := len(str) - 1
	j := 0

	for j < l {
		if j == l {
			continue
		}
		if str[j] != str[l] {
			return false
		}
		j++
		l--
	}
	return true
}

func FindLargestPal() int {
	var a, b int
	maxPal := -1

	for a = 999; a >= 100; a-- {
		for b = a; b >= 100; b-- {
			prod := a * b
			if isPalindrome(prod) {
				if maxPal < prod {
					fmt.Printf("Found Plaindrome is: %d, numbers are (%d, %d)\n", prod, a, b)
					maxPal = prod
				}
			}
		}
	}
	return maxPal
}
