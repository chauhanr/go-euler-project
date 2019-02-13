package levelone

import "testing"

func TestCalcNthPerm(t *testing.T) {
	e := "2783915460"
	p := CalculateNthPerm(1000000)
	if p != e {
		t.Errorf("Expected value of 1 millionth item %s but got %s\n", e, p)
	}
}
