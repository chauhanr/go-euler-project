package leveltwo

import "testing"

func TestConvertBase(t *testing.T) {
	tcases := []struct {
		n    int64
		base int
		r    string
	}{
		{585, 2, "1001001001"},
		{4, 2, "100"},
		{5, 2, "101"},
		{7, 2, "111"},
		{2, 2, "10"},
		{585, 10, "585"},
	}

	for _, tc := range tcases {
		r := convertBase(tc.n, tc.base)
		if r != tc.r {
			t.Errorf("Expected base 2 of %d to be %s but got %s\n", tc.n, tc.r, r)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tcases := []struct {
		s     string
		isPal bool
	}{
		{"585", true},
		{"1001001001", true},
		{"9999", true},
		{"8990", false},
		{"789987", true},
		{"4", true},
		{"100", false},
		{"5", true},
		{"101", true},
		{"111", true},
	}

	for _, tc := range tcases {
		b := isPalindrome(tc.s)
		if b != tc.isPal {
			t.Errorf("Expected %s to be palindrome %t but got %t", tc.s, tc.isPal, b)
		}
	}

}

func TestFindPalSum(t *testing.T) {
	var n, e int64
	n = 1000000
	e = 1999898

	sum := FindPalindromeSum(n)
	if e != sum {
		t.Errorf("Sum of palindrome in two bases is expected to be %d but got %d", e, sum)
	}
}
