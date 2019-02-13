package levelone

import "testing"

func TestCollatzCount(t *testing.T) {
	c13 := CollatzCount(13)
	if c13 != 10 {
		t.Errorf("Expected Collatz count for %d to be %d and got %d", 13, 10, c13)
	}
}

func TestLargestCollatzSeq(t *testing.T) {
	m1 := FindCollatzSeqCount(1000000)
	expected := 525
	if m1 != expected {
		t.Errorf("Expected collatz value for 1m limit is %d but for %d", 10, m1)
	}

}
