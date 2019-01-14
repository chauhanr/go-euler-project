package levelone

import "fmt"

func MaxProduct(grid [][]int, tdim int, l int) int64 {
	var mprod int64
	mprod = -1

	for x := 0; x <= tdim-l; x++ {
		for y := 0; y <= tdim-l; y++ {
			val := maxSubMetric(grid, x, y, y, x, l)
			if val > mprod {
				fmt.Printf("for (%d, %d) and offsets (%d, %d) val is %d\n", x, y, y, x, val)
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
				//	fmt.Printf("value (%d, %d) - offsets (%d,%d) value %d\n", i, j, xoff, yoff, grid[i][j])
				dprod = dprod * int64(grid[i][j])
			}
		}
	}
	if mprod < dprod {
		//fmt.Printf("dprod value max at diagonal (%d, %d) -  %d\n", x, y, dprod)
		mprod = dprod
	}
	for i = x; i <= x+l-1; i++ {
		hprod = 1
		for j = y; j <= y+l-1; j++ {
			hprod = hprod * int64(grid[i][j])
			j++
		}
		if hprod > mprod {
			fmt.Printf("hprod value max at diagonal (%d, %d) -  %d\n", x, y, hprod)
			mprod = hprod
		}
	}

	for j = y; j <= y+l-1; j++ {
		vprod = 1
		for i = x; i <= x+l-1; i++ {
			vprod = vprod * int64(grid[i][j])
		}
		if vprod > mprod {
			fmt.Printf("vprod value max at diagonal (%d, %d) -  %d\n", x, y, vprod)
			mprod = vprod
		}
	}
	//fmt.Printf("hprod %d dprod %d vprod %d\n", hprod, dprod, vprod)
	return mprod
}
