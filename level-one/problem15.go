package levelone

var lgrid [21][21]int64

func initGrid() {
	for i := 0; i < 21; i++ {
		for j := 0; j < 21; j++ {
			lgrid[i][j] = 0
		}
	}
}

func LatticePathCount(dim int64) int64 {
	initGrid()
	paths := travel(0, 0, dim)
	return paths
}

func travel(x, y, l int64) int64 {
	if x < l && y < l {
		if lgrid[x+1][y] == 0 {
			lgrid[x+1][y] = travel(x+1, y, l)
		}
		if lgrid[x][y+1] == 0 {
			lgrid[x][y+1] = travel(x, y+1, l)
		}
		return lgrid[x+1][y] + lgrid[x][y+1]
	} else if x == l && y < l {
		if lgrid[l][y+1] == 0 {
			lgrid[l][y+1] = travel(l, y+1, l)
		}
		return lgrid[l][y+1]
	} else if x < l && y == l {
		if lgrid[x+1][l] == 0 {
			lgrid[x+1][l] = travel(x+1, l, l)
		}
		return lgrid[x+1][y]
	} else {
		return int64(1)
	}
}
