package levelone

import "testing"

func TestFactorialDigitSum(t *testing.T) {
	testcases := []struct {
		n        int
		fact     []int
		sumDigit int
	}{
		{5, []int{1, 2, 0}, 3},
		{10, []int{3, 6, 2, 8, 8, 0, 0}, 27},
		{100, []int{1}, 648},
	}

	for _, tcase := range testcases {
		s := FactorialDigitSum(tcase.n)
		if s != tcase.sumDigit {
			t.Errorf("Expected the sum to be %d but got %d", tcase.sumDigit, s)
		}
	}
}
