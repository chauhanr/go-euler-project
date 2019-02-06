package leveltwo

import (
	"fmt"
	"testing"
)

func TestPopulateBox(t *testing.T) {
	const l = 5
	box := make([][]int, l)
	for i := range box {
		box[i] = make([]int, l)
	}

	err := populateBox(box, l)
	if err != nil {
		t.Errorf("Case failed with message %s", err.Error())
	} else {
		printBox(box, l)
	}

}

func TestCalculateDSum(t *testing.T) {
	tcases := []struct {
		l    int
		dSum int
	}{
		{5, 101},
		{1001, 669171001},
	}

	for _, tc := range tcases {
		box := make([][]int, tc.l)
		for i := range box {
			box[i] = make([]int, tc.l)
		}

		d, err := CalculateDiagonalSum(box, tc.l)
		if err != nil {
			t.Errorf("Following error cases the test case to fail : %s", err.Error())
		} else if d != tc.dSum {
			t.Errorf("Sum for box with dim %d expected to be %d but got %d", tc.l, tc.dSum, d)
		}
	}
}

func printBox(box [][]int, l int) {
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			fmt.Printf(" %2d ", box[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
