package levelone

import "testing"

func TestPyTripletCalc(t *testing.T) {
	prod := 31875000
	p := PyTripletCalc()
	if p != prod {
		t.Errorf("abc expected to be %d but got %d\n", prod, p)
	}

}
