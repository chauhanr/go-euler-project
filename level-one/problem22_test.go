package levelone

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestStrCompare(t *testing.T) {
	namePairs := []struct {
		a string
		b string
		r int
	}{
		{"COLIN", "COLINS", -1},
		{"ABE", "BEN", -1},
		{"BEN", "ABE", 1},
		{"BEN", "BEN", 0},
		{"ZACH", "XMEN", 1},
		{"BEN", "BOON", -1},
		{"BUSTER", "BEN", 1},
		{"BUSTER", "BOON", 1},
	}

	for _, c := range namePairs {
		r := StrCompare(c.a, c.b)
		if r != c.r {
			t.Errorf("For pair (%s, %s) expected the result to be %d but got %d", c.a, c.b, c.r, r)
		}
	}
}

func TestSortName(t *testing.T) {
	names := []string{"ABE", "BUSTER", "COLINS", "CABE", "ANTMAN", "ZACH", "JIM", "BEN", "BOON", "ZAP"}

	for _, n := range names {
		SortName(n)
	}
	t.Logf("Sorted Map %v\n", nameMap)

}

func TestCalcScore(t *testing.T) {
	n := []string{"COLIN"}
	v := 53
	e := calculateScore(n)
	if v != e {
		t.Errorf("Expected value %d but got %d", v, e)
	}
}

func TestNameScores(t *testing.T) {
	dat, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		t.Errorf("Error Reading input file %s", err.Error())
	}
	s := string(dat)
	s = s[1:]
	s = s[:len(s)-1]
	t.Logf("Length of string from file %d", len(s))
	namesList := strings.Split(s, "\",\"")
	if namesList != nil && len(namesList) == 0 {
		t.Errorf("Error cannot split the names")
	}
	t.Logf("Number of names %d", len(namesList))

	CalculateNameScores(namesList)
}
