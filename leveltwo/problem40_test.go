package leveltwo

import "testing"

func TestProdDigit(t *testing.T) {
	var e int64
	e = 210
	p := CalProductDigit()

	if p != e {
		t.Errorf("Product expected to be %d but got %d", e, p)
	}

}
