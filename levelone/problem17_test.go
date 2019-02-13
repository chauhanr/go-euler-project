package levelone

import "testing"

func TestNumberStringify(t *testing.T) {
	tcases := []struct {
		num   int
		value string
		count int
	}{
		{99, "ninety nine", 10},
		{101, "one hundred and one", 16},
		{990, "nine hundred and ninety", 20},
		{111, "one hundred and eleven", 19},
		{80, "eighty", 6},
		{1000, "one thousand", 11},
		{100, "one hundred", 10},
		{1, "one", 3},
		{20, "twenty", 6},
		{700, "seven hundred", 12},
		{879, "eight hundred and seventy nine", 26},
	}

	for _, tcase := range tcases {
		str, c := Stringify(tcase.num)
		if str != tcase.value {
			t.Errorf("Expected string %s but got %s for %d", tcase.value, str, tcase.num)
		}
		if c != tcase.count {
			t.Errorf("for %d expected the char count to be %d but got %d\n", tcase.num, tcase.count, c)
		}

	}
}

func TestSumNumberLength(t *testing.T) {
	eSum := 21124
	s := NumberToWords(1000)
	if s != eSum {
		t.Errorf("Sum of words till 1000 expected to be  %d but got %d", eSum, s)
	}

	s = NumberToWords(99)
	if s != 854 {
		t.Errorf("Sum of words till 99 expected to be  %d but got %d", 854, s)
	}
}
