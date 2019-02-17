package leveltwo

import "testing"

func TestCircularPrimes(t *testing.T) {
	tcases := []struct {
		u      int
		primes int
	}{
		{10, 4},
		{100, 13},
		{1000, 25},
		{1000000, 55},
	}

	for _, tc := range tcases {
		c := GetAllCircularPrimesUnder(tc.u)
		if c != tc.primes {
			t.Errorf("Circular primes under %d expected to be %d, but got %d", tc.u, tc.primes, c)
		}
	}
}

func TestRotate(t *testing.T) {
	tcases := []struct {
		arr []int64
		res []int64
	}{
		{[]int64{1, 3, 3}, []int64{3, 3, 1}},
		{[]int64{3, 3, 1}, []int64{3, 1, 3}},
		{[]int64{1, 1, 7, 9}, []int64{1, 7, 9, 1}},
		{[]int64{7, 3, 9, 7, 1, 1}, []int64{3, 9, 7, 1, 1, 7}},
	}

	for _, tc := range tcases {
		r := rotate(tc.arr)
		t.Logf("rotation of %v expected %v got  %v", tc.arr, tc.res, r)
	}
}

func TestToInt(t *testing.T) {
	tcases := []struct {
		arr []int64
		res int64
	}{
		{[]int64{1, 3, 3}, 133},
		{[]int64{3, 3, 1}, 331},
		{[]int64{1, 1, 7, 9}, 1179},
		{[]int64{7, 3, 9, 7, 1, 1}, 739711},
	}

	for _, tc := range tcases {
		r := toInt(tc.arr)
		if r != tc.res {
			t.Errorf("expected value for %v is %d but got %d", tc.arr, tc.res, r)
		}
	}

}
