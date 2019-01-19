package levelone

import "testing"

func TestSumPowers(t *testing.T) {
	cases := []struct {
		pow int
		sum int
	}{
		{7, 11},
		{8, 13},
		{9, 8},
		{10, 7},
		{1000, 1366},
	}

	for _, tcase := range cases {
		s := SumofPowers(tcase.pow)
		if s != tcase.sum {
			t.Errorf("Expected the sum for power %d to be %d but got %d", tcase.pow, tcase.sum, s)
		}
	}
}
