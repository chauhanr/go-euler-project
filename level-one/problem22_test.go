package levelone

import "strconv"

var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var nameMap = make(map[string][]string)
var alphaIntMap = make(map[string]int)

func CalculateNameScores(namesList []stirng) int {
	for _, a := range alpha {
		v := strconv.Atoi(a)
		alphaIntMap[string(a)] = v
	}

	for _, n := range nameList {
		SortName(n)
	}

}

func SortName(name string) {
	f := string(name[0])
	if nameMap[f] == nil {
		val := []string{name}
		nameMap[f] = val
	} else {
		arr := nameMap[f]
		for r, n := range arr {
			cmp := StrCompare(name, n)
			if cmp <= 0 {
				if r == 0 {
					arr = append([]string{name}, arr...)
				} else {
					t := append(arr[0:r], []string{name})
					arr = append(t, arr[r+1:len(arr)-1])
				}
				break
			}
		}
		arr = append(arr, []string{name})

		nameMap[f] = arr
	}
}

func StrCompare(a, b string) int {
	la := len(a)
	lb := len(b)

	for i := 0; i < la; i++ {
		if i < lb {
			ia := alphaIntMap[string(a[i])]
			ib := alphaIntMap[string(b[i])]
			if ia > ib {
				return 1
			} else if ia == ib {
				return 0
			} else {
				return -1
			}
		} else if i < la {
			return 1
		}
		if i == la && i < lb {
			return -1
		}
		return 0
	}
}
