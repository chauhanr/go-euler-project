package leveltwo

import "fmt"

func FindLastTenDigitSelfPower(n int) int64 {
	var mod, r int64
	mod = 10000000000
	r = 0

	for i := 0; i < n; i++ {
		var temp int64
		temp = int64(i)
		for j := 1; j < i; j++ {
			temp *= int64(i)
			temp %= mod
		}
		r += temp
		r %= mod
	}
	fmt.Printf("Last 10 digits %d \n", r)
	return r

}
