package levelone

import "testing"

var testCases = []struct {
	last           int
	cDiff          int
	divisibleValue int
	nthTerm        int
}{
	{30, 3, 27, 9},
	{30, 5, 25, 5},
	{30, 15, 15, 1},
}

var testCaseSum = []struct {
	v1   int
	v2   int
	last int
	sum  int
}{
	{3, 5, 30, 195},
	{3, 5, 1000, 233168},
}

func TestLastTermFunc(t *testing.T) {
	for _, tcase := range testCases {
		lValue, err := determineLastTerm(tcase.last, tcase.cDiff)
		if err != nil {
			t.Errorf("expected value %d but got error %s", tcase.divisibleValue, err.Error())
		} else if lValue != tcase.divisibleValue {
			t.Errorf("Expected value %d, but got value %d", tcase.divisibleValue, lValue)
		}
	}
}

func TestNthTermFunc(t *testing.T) {
	for _, tcase := range testCases {
		nTerm := determineNthTerm(tcase.cDiff, tcase.cDiff, tcase.divisibleValue)
		if nTerm != tcase.nthTerm {
			t.Errorf("Expected nthTerm as %d, but found %d", tcase.nthTerm, nTerm)
		}
	}
}

func TestSumValues(t *testing.T) {
	for _, tcase := range testCaseSum {
		sum := determineSum(tcase.v1, tcase.v2, tcase.last)
		if sum != tcase.sum {
			t.Errorf("Sum expected to be %d but got calculated sum %d", tcase.sum, sum)
		}
	}
}

func TestProgressionSum(t *testing.T) {
	testCasesPsum := []struct {
		a   int
		n   int
		sum int
	}{
		{3, 9, 135},
		{5, 5, 75},
		{15, 1, 15},
		{3, 333, 166833},
		{5, 199, 99500},
		{15, 66, 33165},
	}

	for _, tcase := range testCasesPsum {
		sum := progressionSum(tcase.a, tcase.a, tcase.n)
		if sum != tcase.sum {
			t.Errorf("Expected sum was %d but got %d", tcase.sum, sum)
		}
	}
}
