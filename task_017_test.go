package mapset

import "testing"

func TestTask017EqualNotProperSuperset(t *testing.T) {
	a := NewThreadUnsafeSet(1, 2)
	b := NewThreadUnsafeSet(2, 1)
	if a.IsProperSuperset(b) {
		t.Fatal("equal set is proper superset")
	}
}
