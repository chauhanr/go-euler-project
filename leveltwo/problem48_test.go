package leveltwo

import "testing"

func TestLastTenDigits(t *testing.T) {
	tcases := []struct {
		n int
		r int64
	}{
		{10, 405071317},
		{1000, 9110846700},
	}

	for _, tc := range tcases {
		r := FindLastTenDigitSelfPower(tc.n)
		if r != tc.r {
			t.Errorf("Expected the last 10 digits for %d to be %d but got %d", tc.n, tc.r, r)
		}
	}

}
