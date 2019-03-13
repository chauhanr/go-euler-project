package leveltwo

import "testing"

func TestGetConvergedTerm(t *testing.T) {
	e := 133333
	v := GetConvergedTerm()
	if v != e {
		t.Errorf("Expceted convered term %d but got %d", e, v)
	}
}

func TestPentagonal(t *testing.T) {
	tcases := []struct {
		term         int
		isPentagonal bool
	}{
		{1, true},
		{9, false},
		{5, true},
		{21, false},
		{22, true},
		{51, true},
		{55, false},
	}

	for _, tc := range tcases {
		isP := isPentagonal(tc.term)
		if isP != tc.isPentagonal {
			t.Errorf("Term %d ispentagonal %t but got %t", tc.term, tc.isPentagonal, isP)
		}

	}
}

func TestHectagonal(t *testing.T) {
	tcases := []struct {
		term   int
		isHect bool
	}{
		{1, true},
		{9, false},
		{6, true},
		{15, true},
		{22, false},
		{28, true},
		{45, true},
	}

	for _, tc := range tcases {
		isH := isHectagonal(tc.term)
		if isH != tc.isHect {
			t.Errorf("Term %d ispentagonal %t but got %t", tc.term, tc.isHect, isH)
		}

	}

}
