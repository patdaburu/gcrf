package subpkg

import "testing"

// TestPlusOne ...
func TestPlusOne(t *testing.T) {
	r := PlusOne(2)
	if r != 3 {
		t.Errorf("PlusOne was incorrect, got: %d, want: %d", r, 3)
	}
}
