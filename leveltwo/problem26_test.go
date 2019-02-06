package leveltwo

import "testing"

func TestFindDecimal(t *testing.T) {
	tCases := []struct {
		d   int
		ln  int
		lim int
	}{
		{2, 0, 16},
		{3, 1, 16},
		{5, 0, 16},
		{6, 1, 16},
		{7, 6, 16},
		{8, 0, 16},
		{9, 1, 16},
		{12, 2, 16},
		{13, 6, 16},
		{17, 16, 32},
		{983, 982, 32},
	}

	for _, tc := range tCases {
		r := FindRecurringDecimalCount(tc.d)
		if r != tc.ln {
			t.Errorf("for num %d expected to have %d recurring digist but got %d elements", tc.d, tc.ln, r)
		}
	}
}

func TestFindLongestRecCount(t *testing.T) {
	max, n := FindLongestDecimalCount(1000)
	if max != 982 && n != 983 {
		t.Errorf("Expected the (max,n) to be (%d,%d) but got (%d,%d)\n", 131, 17, max, n)
	}
}
