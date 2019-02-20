package leveltwo

import "testing"

func TestLargestPan(t *testing.T) {
	p := FindLargestPan()
	var e int64
	e = 932718654

	if e != p {
		t.Errorf("Max Pandigital expected to be %d but got %d", e, p)
	}
}
