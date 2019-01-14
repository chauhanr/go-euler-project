package levelone

import "fmt"

func MaxProduct(grid [][]int, tdim int, l int) int64 {
	var mprod int64
	mprod = -1

	for x := 0; x <= tdim-l; x++ {
		for y := 0; y <= tdim-l; y++ {
			val := maxSubMetric(grid, x, y, x, y, l)
			if val > mprod {
				fmt.Printf("for (%d, %d) and offsets (%d, %d) val is %d\n", x, y, x, y, val)
				mprod = val
			}
		}
	}

	return mprod
}

func maxSubMetric(grid [][]int, x, y, xoff, yoff int, l int) int64 {
	var dprod, hprod, mprod, vprod int64
	dprod = 1
	hprod = 1
	vprod = 1
	mprod = 1

	i := x
	j := y
	for i := x; i <= x+l-1; i++ {
		for j := y; j <= y+l-1; j++ {
			if i+xoff == j+yoff {
				//fmt.Printf("value (%d, %d) value %d\n", i, j, grid[i][j])
				dprod = dprod * int64(grid[i][j])
			}
		}
	}
	j = y
	for j <= y+l-1 {
		hprod = hprod * int64(grid[x][j])
		j++
	}
	i = x
	for i <= x+l-1 {
		vprod = vprod * int64(grid[i][y])
		i++
	}
	//fmt.Printf("hprod %d dprod %d vprod %d\n", hprod, dprod, vprod)
	if hprod > vprod {
		mprod = hprod
	} else {
		mprod = vprod
	}
	if mprod < dprod {
		mprod = dprod
	}
	return mprod
}
