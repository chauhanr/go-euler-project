package leveltwo

import "testing"

func TestGetConsPrimeFactors(t *testing.T) {
	e := 977
	v := GetConsectivePrimeFactorTerm(2)

	if e != v {
		t.Errorf("Expected the first term to be %d but got %v", e, v)
	}
}
