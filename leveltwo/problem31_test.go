package leveltwo

import "testing"

func TestCoinDist(t *testing.T) {
	d := 13000
	p := CalculateCoinPerm([]int{1, 2, 5, 10, 20, 50, 100, 200})
	if p != d {
		t.Errorf("Expected the permutation to be %d but got %d", d, p)
	}
}
