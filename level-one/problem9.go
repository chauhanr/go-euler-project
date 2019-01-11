package levelone

import "fmt"

func PyTripletCalc() int {
	var a, b, c int

	for c = 998; c > 250; c-- {
		for b = 1; b+c < 1000 && c > b; b++ {
			a = 1000 - (b + c)
			if a > b {
				continue
			}
			as := a * a
			bs := b * b
			cs := c * c

			if a+b+c == 1000 && (as+bs) == cs {
				fmt.Printf("Found values (%d, %d, %d)\n", a, b, c)
				return a * b * c
			}
		}
	}
	return -1
}
