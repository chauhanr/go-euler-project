package leveltwo

import (
	"fmt"
	"strconv"
)

func FindLargestPan() int64 {
	var n, maxp, p int64
	maxp = 0
	n = 9876
	isPan := false
	p = 0

	for n >= 4000 {
		if n%10 == 0 {
			isPan, p = BuildPandigital(n / 10)
		} else if n%100 == 0 {
			isPan, p = BuildPandigital(n / 100)
		} else if n%1000 == 0 {
			isPan, p = BuildPandigital(n / 1000)
		} else {
			isPan, p = BuildPandigital(n)
		}
		if isPan && p > maxp {
			//fmt.Printf("prev maxp %d and new maxp: %d\n", maxp)
			maxp = p
		}
		n--
	}
	return maxp
}

func BuildPandigital(n int64) (bool, int64) {
	str := ""
	isPan := true
	m := map[int64]int{}
	if isValidCandidate(n) {
		num := n
		i := 1
		for isPan && !isSetComplete(m) {
			num = n * int64(i)
			isPan = isPandigitalCandidate(m, num)
			if len(str+strconv.FormatInt(num, 10)) <= 9 {
				str += strconv.FormatInt(num, 10)
			}
			i++
		}
		if isSetComplete(m) {
			//fmt.Printf("m : %v for %d\n", m, n)
			isPan = true
		} else {
			isPan = false
		}
	}
	d, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false, 0
	}
	if isPan {
		fmt.Printf("Pandigital for %d is %d found \n", n, d)
	}
	return isPan, d
}

func isSetComplete(m map[int64]int) bool {
	var i int64
	for i = 1; i <= 9; i++ {
		if _, ok := m[i]; !ok {
			return false
		}
	}
	return true
}

func isPandigitalCandidate(m map[int64]int, n int64) bool {
	for n != 0 {
		d := n % 10
		if d == 0 {
			return false
		}
		if _, ok := m[d]; ok {
			return false
		} else {
			m[d] = 1
		}
		n = n / 10
	}
	return true
}

func isValidCandidate(n int64) bool {
	r := true
	if hasZeros(strconv.FormatInt(n, 10)) || (n%5 == 0) {
		r = false
	}
	if hasDigitRepeated(n) {
		r = false
	}
	return r
}

func hasDigitRepeated(n int64) bool {
	m := map[int64]int{}
	for n != 0 {
		d := n % 10
		if _, ok := m[d]; ok {
			return true
		} else {
			m[d] = 1
		}
		n = n / 10
	}
	return false
}

func hasZeros(str string) bool {
	z := "0"
	for _, s := range str {
		if string(s) == z {
			return true
		}
	}
	return false
}
