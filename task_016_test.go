package mapset

import "testing"

func TestTask016EqualNotProperSubset(t *testing.T) {
	a := NewThreadUnsafeSet(1, 2)
	b := NewThreadUnsafeSet(2, 1)
	if a.IsProperSubset(b) {
		t.Fatal("equal set is proper subset")
	}
}
func TestTask016StrictSubset(t *testing.T) {
	a := NewThreadUnsafeSet(1)
	b := NewThreadUnsafeSet(1, 2)
	if !a.IsProperSubset(b) {
		t.Fatal("strict subset rejected")
	}
}
