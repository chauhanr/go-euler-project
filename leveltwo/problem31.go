package leveltwo

import "fmt"

func CalculateCoinPerm(coins []int) int {
	target := 200
	ways := make([]int, target+1)
	ways[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= target; j++ {
			ways[j] += ways[j-coins[i]]
		}
	}
	fmt.Printf("ways %v\n", ways)
	return ways[target]
}
