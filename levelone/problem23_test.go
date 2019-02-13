package levelone

import "testing"

func TestNonAbundantSum(t *testing.T) {
	eSum := 4179871
	ab := FindAbundantSums()
	if ab != eSum {
		t.Errorf("Non Abundant sum expected to be %d but got %d\n", eSum, ab)
	}
}
