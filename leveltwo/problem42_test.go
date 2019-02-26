package leveltwo

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestTermCount(t *testing.T) {
	tcases := []struct {
		w string
		c int
	}{
		{"YOUTH", 89},
		{"TEMPERATURE", 142},
		{"SENTENCE", 85},
	}

	for _, tc := range tcases {
		c := GetTermCount(tc.w)
		if c != tc.c {
			t.Errorf("Expected word count for %s to be %d but for %d", tc.w, tc.c, c)
		}
	}

}

func TestTriangleWords(t *testing.T) {
	dat, err := ioutil.ReadFile("p042_words.txt")
	if err != nil {
		t.Errorf("Error reading the file p042_words.txt %s", err.Error())
	}
	s := string(dat)
	s = s[1:]
	s = s[:len(s)-1]

	wordList := strings.Split(s, "\",\"")
	if wordList == nil || len(wordList) == 0 {
		t.Errorf("Error Spliting the file content")
	}

	tr := FindTriangleTerms(wordList)
	e := 5061
	if tr != 5061 {
		t.Errorf("Expected to find number of triangle words %d but got %d", e, tr)
	}
}
