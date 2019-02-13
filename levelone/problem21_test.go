package levelone

import "testing"

func TestAmicablePairs(t *testing.T) {
	s := FindAmicablePairsSum(10000)

	if s != 31626 {
		t.Errorf("Amicable pairs sum for 10000 expected to be %d, but got %d", 31626, s)
	}
}
