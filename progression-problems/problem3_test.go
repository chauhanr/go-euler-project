package levelone

import (
	"testing"
)

func TestLargetKValue(t *testing.T) {
	testCases := []struct {
		value uint64
		k     uint64
	}{
		{50, 4},
		{91, 7},
		{97, 8},
		{103, 8},
	}

	for _, tcase := range testCases {
		k := determineLargestK(tcase.value)
		if k != tcase.k {
			t.Errorf("Expected value %d but got value of k as %d", tcase.k, k)
		}
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		value   uint64
		isPrime bool
	}{
		{100, false},
		{13, true},
		{21, false},
		{37, true},
		{17, true},
	}

	for _, tcase := range testCases {
		isPrime := isPrime(tcase.value)
		if isPrime != tcase.isPrime {
			t.Errorf("%d value has a mismatch is prime: %t but was found to be %t", tcase.value, tcase.isPrime, isPrime)
		}
	}
}

func TestLargestFactor(t *testing.T) {

	testCases := []struct {
		value       uint64
		primeFactor uint64
	}{
		{100, 5},
		{91, 13},
		{21, 7},
		{35, 7},
		{32, 2},
		{24, 3},
		{9, 3},
		{10, 5},
		{6, 3},
		{5, 5},
		{37, 37},
		{74, 37},
		{7, 7},
		//	{600851475143, 3},
	}

	for _, tcase := range testCases {
		pf := LargestPrimeFactor(tcase.value)
		if pf != tcase.primeFactor {
			t.Errorf("Value : %d Expected largest factor is %d but got %d", tcase.value, tcase.primeFactor, pf)
		}
	}
}

func TestLargestFactor2(t *testing.T) {

	testCases := []struct {
		value       int64
		primeFactor int64
	}{
		{100, 5},
		{91, 13},
		{21, 7},
		{35, 7},
		{32, 2},
		{24, 3},
		{9, 3},
		{10, 5},
		{6, 3},
		{5, 5},
		{37, 37},
		{74, 37},
		{7, 7},
		{600851475143, 6857},
	}

	for _, tcase := range testCases {
		pf := LargestPrimeFactor2(tcase.value)
		if pf != tcase.primeFactor {
			t.Errorf("Value : %d Expected largest factor is %d but got %d", tcase.value, tcase.primeFactor, pf)
		}
	}
}
