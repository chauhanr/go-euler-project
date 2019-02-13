package levelone

import "testing"

func TestSmallestMultiple(t *testing.T) {
	testCases := []struct {
		limit    int64
		multiple int64
	}{
		{10, 2520},
		{15, 360360},
		{20, 232792560},
		{25, 26771144400},
	}

	for _, tcase := range testCases {
		mul := FindSmallestMultiple(tcase.limit)

		if mul != tcase.multiple {
			t.Errorf("Expected the multiple to be %d but found %d\n", tcase.multiple, mul)
		}
	}
}
