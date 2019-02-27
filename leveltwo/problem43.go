package leveltwo

import (
	"fmt"
	"strconv"
)

var digitStr = []string{"4", "8", "7", "6", "5", "9", "3", "2", "1", "0"}
var divisorMap = map[int]int{
	7: 17,
	6: 13,
	5: 11,
	4: 7,
	3: 5,
	2: 3,
	1: 2,
}

func FindSPandigitalSum() int64 {

	sub := map[string]int{}
	permutePandigital(sub, digitStr, 0, 9)
	fmt.Printf("permutations : %v\n", sub)
	var sum int64
	sum = 0
	for key, _ := range sub {
		n, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing number %s error: %s\n", key, err)
		} else {
			sum += n
		}
	}
	fmt.Printf("pandigital sum : %d\n", sum)
	return sum
}

func permutePandigital(sub map[string]int, num []string, l, r int) {
	if num[0] != "0" {
		s := 7
		e := 10
		isDiv := true
		for s != 0 && isDiv {
			n := num[s:e]
			//fmt.Printf("Num: %v n (%d,%d) is : %s\n", num, s, e, stringifyStr(n))
			if !isDivisible(stringifyStr(n), divisorMap[s]) {
				isDiv = false
			}
			s--
			e--
		}
		if isDiv {
			ns := stringifyStr(num)
			if _, ok := sub[ns]; !ok {
				//fmt.Printf("Pandigital %s\n", ns)
				sub[ns] = 1
			}
		}
	}
	if l == r {
		return
	}
	for i := l; i <= r; i++ {
		swapStr(num, l, i)
		permutePandigital(sub, num, l+1, r)
		swapStr(num, l, i)
	}
}

func stringifyStr(str []string) string {
	fStr := ""
	for _, s := range str {
		fStr += s
	}
	return fStr
}

func isDivisible(num string, d int) bool {
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Printf("error parsing num: %s, %s\n", num, err.Error())
		return false
	}
	if n%d == 0 {
		return true
	}
	return false
}

func swapStr(num []string, l, r int) {
	t := num[l]
	num[l] = num[r]
	num[r] = t
}
