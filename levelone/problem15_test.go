package levelone

import "testing"

func TestLatticePath(t *testing.T) {
	testCases := []struct {
		dim   int64
		paths int64
	}{
		{1, 2},
		{2, 6},
		{3, 20},
		{4, 70},
		{5, 252},
		{6, 924},
		{7, 3432},
		{10, 184756},
		{15, 155117520},
		{20, 137846528820},
	}

	for _, tcase := range testCases {
		p := LatticePathCount(tcase.dim)
		if p != tcase.paths {
			t.Errorf("Expected path count for lattice %d is  %d but got %d", tcase.dim, tcase.paths, p)
		}
	}
}
