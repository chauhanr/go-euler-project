package levelone

import (
	"fmt"
	"strconv"
)

func initialize(a []int) []int {
	for i := 0; i < len(a); i++ {
		a[i] = 0
	}
	return a
}

func NthFibSeriesElement(n int) int {
	f1 := []int{1}
	f2 := []int{1}
	fn := []int{0}
	s := 2

	for len(fn) < n {
		fn = nextSeriesElement(f1, f2)
		f1 = f2
		f2 = fn
		s++
	}
	fmt.Printf("Element is %s\n", stringify(fn))
	return s
}

func nextSeriesElement(a, b []int) []int {
	l := 0
	if len(a) > len(b) {
		l = len(a)
		b = padNum(b, l)
	} else if len(a) < len(b) {
		l = len(b)
		a = padNum(a, l)
	} else if len(a) == len(b) {
		l = len(a)
	}
	n := make([]int, l)
	initialize(n)
	r := 0
	//fmt.Printf("a: %v, b: %v\n", a, b)
	for i := l - 1; i >= 0; i-- {
		e := a[i] + b[i] + r
		n[i] = e % 10
		r = e / 10
	}
	for r != 0 {
		n = append([]int{r % 10}, n...)
		r = r / 10
	}
	return n
}

func padNum(a []int, l int) []int {
	s := l - len(a) - 1
	for s >= 0 {
		a = append([]int{0}, a...)
		s--
	}
	return a
}

func stringify(a []int) string {
	s := ""
	for _, e := range a {
		s += strconv.Itoa(e)
	}
	return s
}
