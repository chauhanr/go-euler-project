package levelone

import "strconv"

var perm = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func CalculateNthPerm(n int) string {
	count := 1
	for count < n {
		n := len(perm)
		i := n - 1
		for perm[i-1] >= perm[i] {
			i--
		}
		j := n
		for perm[j-1] <= perm[i-1] {
			j--
		}
		swap(i-1, j-1)
		i++
		j = n
		for i < j {
			swap(i-1, j-1)
			i++
			j--
		}
		count++
	}
	p := ""
	for _, i := range perm {
		p += strconv.Itoa(i)
	}
	return p
}

func swap(i, j int) {
	k := perm[j]
	perm[j] = perm[i]
	perm[i] = k
}
