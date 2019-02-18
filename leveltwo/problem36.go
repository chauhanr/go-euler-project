package leveltwo

import (
	"fmt"
	"strconv"
)

func FindPalindromeSum(n int64) int64 {
	var sum int64
	sum = 0
	var i int64
	for i = 1; i < n; i++ {
		b10 := convertBase(i, 10)
		if isPalindrome(b10) {
			b2 := convertBase(i, 2)
			if isPalindrome(b2) {
				sum += i
			}
		}
	}
	fmt.Printf("Sum for all palindromes in base 10 and 2 less then %d is %d\n", n, sum)
	return sum
}

func isPalindrome(s string) bool {
	l := len(s)
	if l == 1 {
		return true
	}
	if l%2 == 0 {
		s1 := s[:(l / 2)]
		s2 := Reverse(s[(l / 2):])
		if s1 == s2 {
			return true
		}
		return false
	} else {
		s1 := s[:(l / 2)]
		s2 := Reverse(s[(l/2)+1:])
		if s1 == s2 {
			return true
		}
		return false
	}
}

func Reverse(str string) (result string) {
	result = ""
	for _, s := range str {
		result = string(s) + result
	}
	return result

}

func convertBase(n int64, base int) string {
	b := strconv.FormatInt(n, base)
	return b
}
