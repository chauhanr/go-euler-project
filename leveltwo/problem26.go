package leveltwo

import (
	"fmt"
	"strconv"
	"strings"
)

func FindDecimal(dem int, lim int) []int {
	var d float64
	d = float64(1) / float64(dem)
	s := fmt.Sprintf("%f", d)
	fmt.Printf("result : %s\n", s)
	dec := strings.Split(s, ".")[1]
	num := int64(0)
	var err error
	if len(dec) > lim {
		num, err = strconv.ParseInt(dec[:lim], 10, 64)
	} else {
		num, err = strconv.ParseInt(dec, 10, 64)
	}
	if err != nil {
		fmt.Printf("Error when parsing int %d is %s\n", dec, err.Error())
		return []int{}
	}

	a := []int{}
	for num > 0 {
		a = append([]int{int(num % 10)}, a...)
		num = num / 10
	}
	fmt.Printf("arr - %v\n", a)
	return a
}
