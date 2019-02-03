package levelone

import "testing"

func TestNthFibSeries(t *testing.T) {
	n := 3
	e := 12
	v := NthFibSeriesElement(n)
	if v != e {
		t.Errorf("Expected index for 3rd element is  %d but got %d", e, v)
	}
	n = 1000
	e = 10001
	v = NthFibSeriesElement(n)
	if v != e {
		t.Errorf("Expected index for 1000th element is %d but fot %d\n", e, v)
	}

}
