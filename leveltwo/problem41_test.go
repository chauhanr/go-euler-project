package leveltwo

import (
	"testing"
)

func TestFindLargestPrimePan(t *testing.T) {
	var e int64
	e = 7652413
	p := FindLargestPrimePan()
	if e != p {
		t.Errorf("Largest prime pandigital expected %d but got %d", e, p)
	}
}
