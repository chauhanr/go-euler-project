package leveltwo

import (
	"fmt"
	"math"
	"strconv"
)

func FindLargestPrimePan() int64 {
	num := []int{7, 6, 5, 4, 3, 2, 1}

	p := map[int64]int{}
	permute(p, num, 0, len(num)-1)
	fmt.Printf("permutes count  %d\n", len(p))

	var max int64
	max = 0
	for k, _ := range p {
		//fmt.Printf("pan - %d\n", k)
		if k > max {
			fmt.Printf("pan prime - %d\n", k)
			max = k
		}
	}
	fmt.Printf("max prime pandigital : %d", max)
	return max
}

func permute(p map[int64]int, num []int, l, r int) {
	n := getNumber(num)
	if _, ok := p[n]; !ok {
		if (n-1)%6 == 0 || (n+1)%6 == 0 {
			if IsPrimeInt(n) {
				p[n] = 1
			}
		}
	}
	if l == r {
		return
	}
	for i := l; i <= r; i++ {
		swap(num, l, i)
		permute(p, num, l+1, r)
		swap(num, l, i)
	}

}

func swap(num []int, l, r int) {
	t := num[l]
	num[l] = num[r]
	num[r] = t
}

func getNumber(num []int) int64 {
	var n int64
	n = 0
	l := len(num) - 1
	for i := l; i >= 0; i-- {
		p := int64(math.Pow10(l - i))
		n += int64(num[i]) * p
	}
	return n
}

func isPandigitalNum(num string) bool {
	fmt.Printf("Checking for pandigital check %s\n", num)
	base := len(num)
	m := map[string]int{}
	for _, r := range num {
		d := string(r)
		if _, ok := m[d]; ok {
			return false
		} else {
			m[d] = 1
		}
	}

	for i := 1; i <= base; i++ {
		k := strconv.Itoa(i)
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}
