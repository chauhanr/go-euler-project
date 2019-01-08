package levelone

import (
	"testing"
)

func TestEvenSumFib(t *testing.T) {
	testCases := []struct {
		limit int64
		sum   int64
	}{
		{100, 44},
		{150, 188},
		{650, 798},
		{4000000, 4613732},
	}

	for _, tcase := range testCases {
		sum := SumEvenFibSeries(tcase.limit)
		if sum != tcase.sum {
			t.Errorf("Expected sum was %d but got %d", tcase.sum, sum)
		}
	}
}
