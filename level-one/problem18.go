package levelone

import "fmt"

var solution = [15][29]int{}

func initSolution() {
	for x := 0; x < 15; x++ {
		for y := 0; y < 29; y++ {
			//fmt.Printf("(%d,%d)\n", x, y)
			solution[x][y] = 0
		}
	}
}

func LargestPathSum(x, y, xn, yn int, py [][]int) int {
	initSolution()
	traversePy(x, y, xn, yn, 0, py)
	printSolution(xn, yn)
	max := 0
	i := xn - 1
	for j := 0; j <= yn-1; j++ {
		if max < solution[i][j] {
			max = solution[i][j]
		}
	}
	return max
}

func printSolution(xn, yn int) {
	for x := 0; x < xn; x++ {
		for y := 0; y < yn; y++ {
			fmt.Printf("%4d ", solution[x][y])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func traversePy(x, y, xn, yn int, pSum int, py [][]int) {
	if x == xn || y < 0 || y == yn {
		return
	}
	//fmt.Printf("(%d,%d)\n", x, y)
	pSum += py[x][y]
	if solution[x][y] <= pSum {
		solution[x][y] = pSum
	} else {
		return
	}
	traversePy(x+1, y-1, xn, yn, pSum, py)
	traversePy(x+1, y+1, xn, yn, pSum, py)
}
