package sheet

import "testing"

func TestToColID(t *testing.T) {
	wanted := []string{"A", "Z", "AA", "AB", "YY", "AAA", "CCC"}
	nums := []int{1, 26, 27, 28, 675, 703, 2109}

	for i, w := range wanted {
		got := intToColID(nums[i])
		if got != w {
			t.Errorf("wanted colID %s, got %s\n", w, got)
		}
	}
}

func TestColIDToInt(t *testing.T) {
	wanted := []int{1, 26, 27, 28, 675, 703, 2109}
	colIDs := []string{"A", "Z", "AA", "AB", "YY", "AAA", "CCC"}

	for i, w := range wanted {
		got := colIDToInt(colIDs[i])
		if got != w {
			t.Errorf("wanted int %d, got %d\n", w, got)
		}
	}
}
