package leveltwo

import "testing"

func TestFindNextPrime(t *testing.T) {
	tcases := []struct {
		n  int
		np int
	}{
		{4, 5},
		{11, 13},
		{10, 11},
		{15, 17},
		{19, 23},
		{31, 37},
	}

	for _, tc := range tcases {
		np := findNextPrime(tc.n)
		if np != tc.np {
			t.Errorf("Expected the next prime after %d to be %d but got %d", tc.n, tc.np, np)
		}

	}
}

func TestLongestQuadValueRun(t *testing.T) {
	tcases := []struct {
		a int
		b int
		r int
	}{
		{1, 41, 40},
		{-79, 1601, 80},
	}

	for _, tc := range tcases {
		l := LongestRunQuadValue(tc.a, tc.b)
		if l != tc.r {
			t.Errorf("Longest run for (%d,%d) was expected to be %d, but got %d ", tc.a, tc.b, tc.r, l)
		}
	}
}

func TestCalcQuadCoefProd(t *testing.T) {

	p := CalculateQuadCoefProd()
	e := -59231
	if p != e {
		t.Errorf("Product of coefficients with highest run is supposed to be %d but got %d", e, p)
	}
}
