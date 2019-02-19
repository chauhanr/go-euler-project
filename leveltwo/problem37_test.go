package leveltwo

import "testing"

func TestLRPrimeSum(t *testing.T) {
	var e int64
	e = 748317
	r := FindLeftRightPrimeSum()
	if r != e {
		t.Errorf("LR sum expected to be %d but got %d", e, r)
	}
}

func TestGetNextPrime(t *testing.T) {
	tcases := []struct {
		n int64
		p int64
	}{
		{9, 11},
		{13, 17},
		{31, 37},
		{81, 83},
		{200, 211},
	}

	for _, tc := range tcases {
		p := GetNextPrime(tc.n)
		if p != tc.p {
			t.Errorf("The next prime to %d is expected to be %d but got %d", tc.n, tc.p, p)
		}
	}
}

func TestLeftRightPrime(t *testing.T) {
	tcases := []struct {
		n  int64
		rp bool
		lp bool
	}{
		{3797, true, true},
		{37, true, true},
		{45, false, false},
	}

	for _, tc := range tcases {
		rp := isRightPrime(tc.n)
		if rp != tc.rp {
			t.Errorf("%d is right prime: %t but found %t", tc.n, tc.rp, rp)
		}
		lp := isLeftPrime(tc.n)
		if lp != tc.lp {
			t.Errorf("%d is left prime: %t but found %t", tc.n, tc.lp, lp)
		}
	}
}
