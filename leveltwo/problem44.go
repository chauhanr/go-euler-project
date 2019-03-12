package leveltwo

import (
	"fmt"
	"math"
)

var penta = []int{}
var m = map[int]int{}

func generatePenagonal(n int) {
	i := 1
	for i <= n {
		term := i * (3*i - 1) / 2
		penta = append(penta, term)
		m[term] = 1
		i++
	}
	//	fmt.Printf("%v\n", m)
}

func GetValueD(n int) int {
	generatePenagonal(n)
	max := penta[len(penta)-1]
	mind := math.MaxInt32

	for i := 0; i <= len(penta)-1; i++ {
		for j := i + 1; j <= len(penta)-1; j++ {
			s := penta[i] + penta[j]
			d := penta[j] - penta[i]

			if s > max {
				break
			}
			_, sok := m[s]
			_, dok := m[d]

			if sok && dok {
				fmt.Printf("i : %d, j : %d\n", i, j)
				if mind > d {
					mind = d
				}
			}
		}
	}
	return mind
}
