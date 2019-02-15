package leveltwo

import "testing"

func TestGetDigitFactorialSum(t *testing.T) {
	tcases := []struct {
		n   int64
		sum int64
	}{
		{145, 145},
		{11, 2},
		{100, 3},
		{333, 18},
		{900, 362882},
	}

	for _, tc := range tcases {
		s := GetDigitFactorialSum(tc.n)
		if s != tc.sum {
			t.Errorf("For number %d expected sum to be %d, but got %d", tc.n, tc.sum, s)
		}
	}
}

func TestFindSumSpecial(t *testing.T) {
	var e int64
	e = 1900
	s := FindAllFactSumN()
	if s != e {
		t.Errorf("Sum of all special number expected %d but got %d", e, s)
	}

}
