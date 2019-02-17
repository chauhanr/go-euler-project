package leveltwo

import (
	"fmt"
	"math"
)

func GetAllCircularPrimesUnder(n int) int {
	cPrimes := map[int64]int{}
	notCPrimes := map[int64]int{}
	cPrimesMap := map[int64][]int64{}
	count := 0
	for i := 2; i < n; i++ {
		if _, ok := cPrimes[int64(i)]; ok {
			continue
		}
		if _, ok := notCPrimes[int64(i)]; ok {
			continue
		}

		if IsPrimeInt(int64(i)) {
			p, isCprime := isCircularPrime(int64(i))
			if isCprime {
				fmt.Printf("Cprimes for n %d,  %v\n", i, p)
				for _, v := range p {
					cPrimes[v] = 1
				}
				cPrimesMap[int64(i)] = p
				count += len(p) + 1
			} else {
				for _, v := range p {
					notCPrimes[v] = 1
				}
			}
		}
	}
	fmt.Printf("Circular Primes : %v\n", cPrimesMap)
	return count
}

func isCircularPrime(n int64) ([]int64, bool) {
	isCircularP := true
	nArr := toArray(n)
	perm := []int64{}
	if len(nArr) == 1 {
		return perm, isCircularP
	}
	permSet := map[int64]int{}
	permSet[n] = 1

	for i := 0; i < len(nArr)-1; i++ {
		nArr = rotate(nArr)
		rn := toInt(nArr)
		if !IsPrimeInt(int64(rn)) {
			isCircularP = false
		}
		if _, ok := permSet[rn]; !ok {
			permSet[rn] = 1
			perm = append(perm, rn)
		}
	}

	return perm, isCircularP
}

func rotate(n []int64) []int64 {
	l := len(n) - 1
	first := n[0]
	for i := 0; i < l; i++ {
		n[i] = n[i+1]
	}
	n[l] = first
	return n
}

func toInt(n []int64) int64 {
	var num int64
	num = 0
	l := len(n) - 1
	for i := l; i >= 0; i-- {
		if i == l {
			num += n[i]
		} else {
			f := math.Pow10(l - i)
			num += int64(f) * n[i]
		}
	}
	return num
}

func toArray(n int64) []int64 {
	dArray := []int64{}
	num := n
	for num != 0 {
		d := int64(num % 10)
		dArray = append([]int64{d}, dArray...)
		num = num / 10
	}
	return dArray
}
