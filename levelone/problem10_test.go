package levelone

import "testing"

func TestSumPrimeBelow(t *testing.T) {
	testCases := []struct {
		limit int64
		sum   int64
	}{
		{10, 17},
		{20, 77},
		{2000000, 142913828922},
	}

	for _, tcase := range testCases {
		psum := SumPrimeBelow(tcase.limit)
		if psum != tcase.sum {
			t.Errorf("Expected the sum to be %d but got %d", tcase.sum, psum)
		}
	}
}
