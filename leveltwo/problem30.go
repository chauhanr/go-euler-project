package leveltwo

import (
	"errors"
	"fmt"
	"math"
)

var dMap map[int]int

func initDmap(p int) {
	dMap = make(map[int]int)
	for i := 2; i <= 9; i++ {
		pow := math.Pow(float64(i), float64(p))
		dMap[i] = int(pow)
	}
	dMap[1] = 1
	dMap[0] = 0
}

func CalculatePowSum(p int) (int, error) {
	nums := []int{}
	initDmap(p)
	if p >= 10 {
		return -1, errors.New("func does not support more than 10 digit number")
	}
	lf := math.Pow(10.0, float64(p+1))
	for i := 2; i < int(lf); i++ {
		s := PowerSum(i, p)
		if s == i {
			nums = append(nums, s)
		}
	}
	sum := 0
	fmt.Printf("Nums : %v\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum, nil
}

func PowerSum(n int, p int) int {
	num := n
	sum := 0
	for num != 0 {
		d := num % 10
		num = num / 10
		if d == 0 {
			sum += d
		} else {
			v := dMap[d]
			sum += v
		}
	}
	return sum
}
