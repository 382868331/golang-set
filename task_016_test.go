package mapset

import "testing"

func TestTask016EqualNotProperSubset(t *testing.T) {
	a := NewThreadUnsafeSet(1, 2)
	b := NewThreadUnsafeSet(2, 1)
	if a.IsProperSubset(b) {
		t.Fatal("equal set is proper subset")
	}
}
