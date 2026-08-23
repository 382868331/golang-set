package mapset

import "testing"

func TestTask020RemovePresent(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2)
	s.Remove(1)
	if s.Contains(1) || s.Cardinality() != 1 {
		t.Fatalf("set=%v", s)
	}
}
