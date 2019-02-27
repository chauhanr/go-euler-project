package leveltwo

import "testing"

func TestFindSPandigitalSum(t *testing.T) {
	var e int64
	e = 16695334890
	i := FindSPandigitalSum()
	if i != e {
		t.Errorf("Sum of pandigital number is %d but got %d", e, i)
	}
}
