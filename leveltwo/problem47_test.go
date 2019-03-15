package leveltwo

import "testing"

func TestGetConsPrimeFactors(t *testing.T) {
	e := 134043
	v := GetConsectivePrimeFactorTerm(4)

	if e != v {
		t.Errorf("Expected the first term to be %d but got %v", e, v)
	}
}
