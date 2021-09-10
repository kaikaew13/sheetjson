package sheet

import "testing"

func TestToColID(t *testing.T) {
	wanted := []string{"A", "Z", "AA", "AB", "YY", "AAA", "CCC"}
	nums := []int{1, 26, 27, 28, 675, 703, 2109}

	for i, n := range nums {
		got := toColID(n)
		if got != wanted[i] {
			t.Errorf("wanted colID %s, got %s\n", wanted[i], got)
		}
	}
}
