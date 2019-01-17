package levelone

import "fmt"

func AddLargeNumber(grid [][]int, d int, r int) []int {
	var res = make([]int, d)
	var rem int
	rem = 0
	var i, j int

	for i = d - 1; i >= 0; i-- {
		sum := rem
		for j = 0; j <= r-1; j++ {
			//fmt.Printf("(%2d, %2d)", j, i)
			sum += grid[j][i]
		}
		ld := sum % 10
		res[i] = ld
		rem = sum / 10
	}

	for rem != 0 {
		d := rem % 10
		sl := []int{d}
		res = append(sl, res...)
		rem = rem / 10
	}
	return res

}

func printGrid(grid [][]int, c, r int) {
	for i := 0; i < r-1; i++ {
		for j := 0; j < c-1; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("\n")
	}
}
