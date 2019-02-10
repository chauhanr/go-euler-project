package leveltwo

import "testing"

func TestPrimeBelow(t *testing.T) {
	tcases := []struct {
		n  int
		pc int64
	}{
		{10, 4},
		{100, 25},
	}

	for _, tc := range tcases {
		pc := primeBelow(tc.n)
		if int64(pc) != tc.pc {
			t.Errorf("Expected prime count below %d to be %d got %d", tc.n, tc.pc, pc)
		}
	}

}
