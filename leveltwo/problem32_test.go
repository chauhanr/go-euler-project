package leveltwo

import "testing"

func TestIsPandigital(t *testing.T) {
	tcases := []struct {
		n   int64
		pan bool
	}{
		{123456789, true},
		{908765412, false},
		{391867254, true},
	}

	for _, tc := range tcases {
		p := IsPandigital(tc.n)
		if p != tc.pan {
			t.Errorf("Number %d is pandigital: %t but got %t", tc.n, tc.pan, p)
		}
	}
}

func TestFindPandigitalSum(t *testing.T) {
	pd := FindPandigitalSum()
	var panSum int64
	panSum = 3121232334
	if pd != panSum {
		t.Errorf("Pandigital sum for 1..9 expected to be %d but got %d", panSum, pd)
	}
}

func TestIsDigitRepeat(t *testing.T) {
	tcases := []struct {
		a int64
		b bool
	}{
		{12, false},
		{111, true},
		{9876543, false},
		{123456789, false},
		{98089, true},
	}

	for _, tc := range tcases {
		b := isDigitRepeated(tc.a)
		if b != tc.b {
			t.Errorf("Number %d has repeated digits : %t but got %t", tc.a, tc.b, b)
		}
	}
}

func TestConcatNumbers(t *testing.T) {
	tcases := []struct {
		a int64
		b int64
		c int64
	}{
		{12, 7254, 127254},
		{7254, 12, 725412},
		{134, 223, 134223},
		{1234, 56789, 123456789},
	}

	for _, tc := range tcases {
		n := concatNumbers(tc.a, tc.b)
		if n != tc.c {
			t.Errorf("For (%d, %d) Expected the number to be %d but got %d", tc.a, tc.b, tc.c, n)
		}
	}
}
