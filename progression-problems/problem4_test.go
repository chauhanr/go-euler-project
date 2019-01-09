package levelone

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCase := []struct {
		value        int
		isPalindrome bool
	}{
		{9999, true},
		{99, true},
		{911, false},
		{99199, true},
		{997899, false},
		{99899, true},
		{11011, true},
	}

	for _, tcase := range testCase {
		ispal := isPalindrome(tcase.value)
		if ispal != tcase.isPalindrome {
			t.Errorf("%d is Palindrome : %t", tcase.value, ispal)
		}
	}
}

func Test3Palindrome(t *testing.T) {
	expected := 906609
	value := FindLargestPal()
	if value != expected {
		t.Errorf("Expected value %d but got %d", expected, value)
	}
}
