package levelone

import "fmt"

func SumofPowers(n int) int {
	p := 0
	var lNum = []int{1}
	for i := 0; i < n; i++ {
		j := len(lNum) - 1
		for ; j >= 0; j-- {
			p = lNum[j]*2 + p
			lNum[j] = p % 10
			p = p / 10
			for j == 0 && p != 0 {
				lNum = append([]int{p % 10}, lNum...)
				p = p / 10
			}
		}
	}
	fmt.Printf("for power %d arr is %v\n", n, lNum)

	sum := 0
	for _, d := range lNum {
		sum += d
	}
	return sum
}
