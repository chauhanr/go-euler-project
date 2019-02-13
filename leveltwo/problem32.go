package leveltwo

import (
	"fmt"
	"strconv"
)

func FindPandigitalSum() int64 {
	var i int64
	var j int64
	var sum int64
	sum = 0
	pMap := map[int64]int64{}

	for i = 2; i < 10; i++ {
		for j = 1234; j <= 10000/(i+1); j++ {
			p := i * j
			c := concatNumbers(i, j, p)
			if IsPandigital(c) {
				if _, ok := pMap[p]; !ok {
					fmt.Printf("(%d * %d) = %d is Pandigital \n", i, j, p)
					sum += p
					pMap[p] = 1
				}
			}
		}
	}
	for i = 11; i < 100; i++ {
		for j = 123; j <= 10000/(i+1); j++ {
			p := i * j
			c := concatNumbers(i, j, p)
			if IsPandigital(c) {
				if _, ok := pMap[p]; !ok {
					fmt.Printf("(%d * %d) = %d is Pandigital \n", i, j, p)
					sum += p
					pMap[p] = 1
				}
			}
		}
	}

	return sum
}

func isDigitRepeated(a int64) bool {
	store := map[int64]int64{}
	for a != 0 {
		//fmt.Printf("Store for %d is %v\n", a, store)
		d := a % 10
		if _, ok := store[int64(d)]; ok {
			return true
		} else {
			store[d] = 1
		}
		a = a / 10
	}
	return false
}

func concatNumbers(a ...int64) int64 {
	ns := ""
	for _, n := range a {
		nstr := strconv.FormatInt(n, 10)
		ns += nstr
	}
	n, err := strconv.ParseInt(ns, 10, 64)
	if err != nil {
		fmt.Printf("Error in parsing integer %s\n", err.Error())
	}
	return n
}

func IsPandigital(p int64) bool {
	dp := []bool{false, false, false, false, false, false, false, false, false, false}
	for p != 0 {
		d := p % 10
		p = p / 10
		if d == 0 {
			return false
		}
		if !dp[d] {
			dp[d] = true
		} else {
			return false
		}
	}
	for i := 1; i < len(dp); i++ {
		if !dp[i] {
			return false
		}
	}
	return true
}
