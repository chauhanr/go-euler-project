package leveltwo

import "testing"

func TestFindDecimal(t *testing.T) {
	tCases := []struct {
		d   int
		ln  int
		lim int
	}{
		{2, 1, 16},
		{3, 16, 16},
		{5, 1, 16},
		{6, 16, 16},
		{7, 16, 16},
		{9, 16, 16},
	}

	for _, tc := range tCases {
		r := FindDecimal(tc.d, tc.lim)
		lr := len(r)
		if lr != tc.ln {
			t.Errorf("Expected the array %v to have %d elements but got %d elements", r, tc.ln, lr)
		}
	}
}
