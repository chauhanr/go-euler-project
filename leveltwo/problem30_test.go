package leveltwo

import (
	"testing"
)

func TestInitDMap(t *testing.T) {
	initDmap(4)
	t.Logf("Init map print for 4 is %v", dMap)
}

func TestPowerSum(t *testing.T) {
	tcases := []struct {
		p int
		s int
	}{
		{4, 19316},
		{3, 1301},
		{5, 443839},
	}

	for _, tc := range tcases {
		s, err := CalculatePowSum(tc.p)
		if err != nil {
			t.Errorf("Method cannot handle the condiitions %s", err.Error())
		}
		if s != tc.s {
			t.Errorf("Expected the sum for power %d to be %d but found %d\n", tc.p, tc.s, s)
		}
	}
}
