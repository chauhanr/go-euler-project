package leveltwo

import (
	"testing"
)

func TestGetValueD(t *testing.T) {
	eD := 32
	aD := GetValueD(2500)

	if eD != aD {
		t.Errorf("Expected value of D is %d but got %d", eD, aD)
	}
}
