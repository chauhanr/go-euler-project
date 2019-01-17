package levelone

import "testing"

func TestSumN(t *testing.T) {
	testcases := []struct {
		n   int64
		sum int64
	}{
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
		{10, 55},
	}

	for _, tcase := range testcases {
		s := SumN(tcase.n)
		if s != tcase.sum {
			t.Errorf("Expected value of sum for %d is %d but ot %d", tcase.n, tcase.sum, s)
		}
	}

}

func TestFactorCombCount(t *testing.T) {
	testcases := []struct {
		sum   int64
		count int64
	}{
		{3, 2},
		{6, 4},
		{28, 6},
		{55, 4},
	}

	for _, tcase := range testcases {
		c, _ := FactorCombCount(tcase.sum)
		if c != tcase.count {
			t.Errorf("Count of multiple for sum %d must be %d but got %d", tcase.sum, tcase.count, c)

		}

	}
}

func TestMultiple500(t *testing.T) {
	c, sum := CalculateMultipleCount()
	var exMul, exSum int64
	exMul = 576
	exSum = 76576500

	if c != exMul || sum != exSum {
		t.Errorf("Expected value Triangle Number %d but got %d", exSum, sum)
	}
}
