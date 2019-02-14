package leveltwo

import "testing"

func TestAreFactionEqual(t *testing.T) {
	tcases := []struct {
		a Fraction
		b Fraction
		r bool
	}{
		{Fraction{49, 98}, Fraction{4, 8}, true},
		{Fraction{4, 8}, Fraction{1, 2}, true},
		{Fraction{1, 3}, Fraction{9, 27}, true},
		{Fraction{1, 2}, Fraction{1, 3}, false},
	}

	for _, tc := range tcases {
		r := AreFractionsEqual(tc.a, tc.b)
		if r != tc.r {
			t.Errorf("Fractions %v and %v expected to be equal %t but got %t", tc.a, tc.b, tc.r, r)
		}
	}

}

func TestRemoveCommonDigit(t *testing.T) {
	tcases := []struct {
		a Fraction
		r Fraction
	}{
		{Fraction{49, 98}, Fraction{4, 8}},
		{Fraction{14, 81}, Fraction{4, 8}},
		{Fraction{42, 31}, Fraction{42, 31}},
		{Fraction{48, 84}, Fraction{4, 4}},
		{Fraction{88, 99}, Fraction{8, 9}},
	}

	for _, tc := range tcases {
		r := removeCommonDigit(tc.a)
		t.Logf("For %v Expected %v got %v", tc.a, tc.r, r)
		if r.num != tc.r.num || r.den != tc.r.den {
			t.Errorf("For Fraction %v expected result to be %v but got %v", tc.a, tc.r, r)
		}
	}
}

func TestFindProductDen(t *testing.T) {
	r := 40
	res := FindProductDen()
	if r != res {
		t.Errorf("Product for demonimator expected to be %d, but got %d", r, res)
	}
}
