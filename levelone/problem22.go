package levelone

import (
	"fmt"
	//"strconv"
)

var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var nameMap = make(map[string][]string)
var alphaIntMap = make(map[string]int)

func initAlphaMap() {
	for i, a := range alpha {
		v := i + 1
		alphaIntMap[string(a)] = v
	}
}

func CalculateNameScores(namesList []string) int {
	if len(alphaIntMap) == 0 {
		initAlphaMap()
	}
	for _, n := range namesList {
		SortName(n)
	}

	f := []string{}
	for _, a := range alpha {
		val := nameMap[string(a)]
		f = append(f, val...)
	}

	score := calculateScore(f)
	fmt.Printf("Score for %d names is :  %d\n", len(f), score)
	return score

}

func calculateScore(sNames []string) int {
	if len(alphaIntMap) == 0 {
		initAlphaMap()
	}
	score := 0
	for i, n := range sNames {
		ns := 0
		for _, a := range n {
			v := alphaIntMap[string(a)]
			ns += v
		}
		score += (i + 1) * ns
	}
	return score
}

func SortName(name string) {
	f := string(name[0])
	if nameMap[f] == nil {
		val := []string{name}
		nameMap[f] = val
	} else {
		placed := false
		arr := nameMap[f]
		for r, n := range arr {
			cmp := StrCompare(name, n)
			if cmp <= 0 {
				placed = true
				if r == 0 {
					arr = append([]string{name}, arr...)
				} else {
					arr = append(arr[0:r], append([]string{name}, arr[r:]...)...)
				}
				break
			}
		}
		if !placed {
			arr = append(arr, []string{name}...)
		}
		nameMap[f] = arr
	}
}

func StrCompare(a, b string) int {
	la := len(a)
	lb := len(b)
	if len(alphaIntMap) == 0 {
		initAlphaMap()
	}
	i := 0
	for ; i < la; i++ {
		if i < lb {
			ia := alphaIntMap[string(a[i])]
			ib := alphaIntMap[string(b[i])]
			if ia > ib {
				return 1
			} else if ia == ib {
				continue
			} else {
				return -1
			}
		} else if i < la {
			return 1
		}
	}
	if i == la && i < lb {
		return -1
	}
	return 0
}
