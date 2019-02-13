package levelone

import "testing"

func TestNthPrimeFunc(t *testing.T) {
	testCases := []struct {
		n     int
		prime int64
	}{
		{1, 2},
		{2, 3},
		{5, 11},
		{17, 59},
		{100, 541},
		{10001, 104743},
	}
	for _, tcase := range testCases {
		nPrime := FindNthPrime(tcase.n)
		if nPrime != tcase.prime {
			t.Errorf("Expected %dth prime to be %d but got %d\n", tcase.n, tcase.prime, nPrime)
		}
	}
}
