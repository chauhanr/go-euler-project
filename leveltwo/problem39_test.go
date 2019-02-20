package leveltwo

import "testing"

func TestFindMaxTripletCount(t *testing.T) {
	p := 1000
	m := FindMaxTrpiletCount(p)
	e := 8

	if m != e {
		t.Errorf("For perimeter %d the max triplet count expected is %d but got %d", p, e, m)
	}
}
