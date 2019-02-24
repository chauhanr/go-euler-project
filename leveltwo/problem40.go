package leveltwo

import (
	"fmt"
	"strconv"
)

func CalProductDigit() int64 {
	var i, index, prod, p int64
	prod = 1
	p = 10
	for i = 1; index <= 1000000; i++ {
		n := strconv.FormatInt(i, 10)
		if index >= p {
			d := getIndexDigit(n, index, p)
			if d == 0 {
				return 0
			}
			prod *= d
			p = nextIndex(p)
		}
		index += int64(len(n))
	}
	fmt.Printf("prod of the digits is :%d\n", prod)
	return prod
}

func getIndexDigit(n string, index, p int64) int64 {
	if index == p {
		r := n[len(n)-1]
		d, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Printf("error parsing digit %s\n", err)
			return 0
		}
		fmt.Printf("n: %s p : %d, index : %d, d: %d\n", n, p, index, d)
		return int64(d)
	} else {
		l := int64(len(n) - 1)
		r := n[l-(index-p)]
		d, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Printf("error parsing digit %s\n", err)
			return 0
		}
		fmt.Printf("n: %s l: %d, p : %d, index : %d, d: %d\n", n, l, p, index, d)
		return int64(d)

	}
}

func nextIndex(p int64) int64 {
	switch p {
	case 10:
		return 100
	case 100:
		return 1000
	case 1000:
		return 10000
	case 10000:
		return 100000
	case 100000:
		return 1000000
	}
	return 0
}
