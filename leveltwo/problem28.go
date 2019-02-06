package leveltwo

import (
	"errors"
	"fmt"
)

const R = 0
const D = 1
const L = 2
const U = 3

var order = []int{R, D, L, U}

func initializeBox(box [][]int, l int) error {
	if l%2 == 0 {
		return errors.New("Box needs to have odd dimention")
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			box[i][j] = 0
		}
	}
	return nil
}

func CalculateDiagonalSum(box [][]int, l int) (int, error) {
	err := populateBox(box, l)
	if err != nil {
		return -1, err
	}
	d1 := 0
	d2 := 0
	i := 0
	j := l - 1
	for i < l && j >= 0 {
		d1 += box[i][i]
		d2 += box[j][i]
		i++
		j--
	}
	return d1 + d2 - 1, nil
}

/*
   this method will populate the box first initialize it and then
   fill in the values starting (l/2, l/2) and moving outward siprally.
*/
func populateBox(box [][]int, l int) error {
	err := initializeBox(box, l)
	if err != nil {
		return err
	}
	v := 1
	mv := -1
	x := (l - 1) / 2
	y := (l - 1) / 2
	box[x][y] = v
	v++
	for v <= l*l {
		mv = nextMove(box, x, y, mv)
		switch mv {
		case R:
			y++
			box[x][y] = v
			break
		case D:
			x++
			box[x][y] = v
			break
		case L:
			y--
			box[x][y] = v
			break
		case U:
			x--
			box[x][y] = v
			break
		default:
			fmt.Printf("do nothing")
		}
		v++
	}
	return nil
}

func nextMove(box [][]int, x, y, cm int) int {
	if cm == -1 {
		return R
	}
	nv := cm
	switch cm {
	case R:
		if x == len(box)-1 {
			nv = -1
		} else if box[x+1][y] == 0 {
			nv = order[R+1] % 4
		} else {
			nv = R
		}
	case D:
		if y == 0 {
			nv = -1
		} else if box[x][y-1] == 0 {
			nv = order[D+1] % 4
		} else {
			nv = D
		}
	case L:
		if x == 0 {
			nv = -1
		} else if box[x-1][y] == 0 {
			nv = order[L+1] % 4
		} else {
			nv = L
		}
	case U:
		if y == len(box)-1 {
			nv = -1
		} else if box[x][y+1] == 0 {
			nv = (U + 1) % 4
		} else {
			nv = U
		}
	}
	return nv
}
