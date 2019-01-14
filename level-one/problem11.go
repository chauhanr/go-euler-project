package levelone

import "fmt"

func MaxProduct(grid [][]int, tdim int, l int) int64 {
	var mprod int64
	mprod = -1
	for x := 0; x <= tdim-l; x++ {
		for y := 0; y <= tdim-l; y++ {
			val := maxSubMetric(grid, x, y, y, x, l)
			if val > mprod {
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
	mprod = -1

	i := x
	j := y
	for i := x; i <= x+l-1; i++ {
		for j := y; j <= y+l-1; j++ {
			fmt.Printf("%2d ", grid[i][j])
			if i+xoff == j+yoff {
				dprod = dprod * int64(grid[i][j])
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
	if mprod < dprod {
		mprod = dprod
	}
	for i = x; i <= x+l-1; i++ {
		hprod = 1
		for j = y; j <= y+l-1; j++ {
			hprod = hprod * int64(grid[i][j])
		}
		if hprod > mprod {
			mprod = hprod
		}
	}

	for j = y; j <= y+l-1; j++ {
		vprod = 1
		for i = x; i <= x+l-1; i++ {
			vprod = vprod * int64(grid[i][j])
		}
		if vprod > mprod {
			mprod = vprod
		}
	}
	fmt.Printf("Max Prod %d\n\n", mprod)
	return mprod
}
