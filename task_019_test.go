package mapset

import "testing"

func TestTask019PopNExact(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3)
	items, n := s.PopN(2)
	if n != 2 || len(items) != 2 || s.Cardinality() != 1 {
		t.Fatalf("count=%d items=%v remaining=%d", n, items, s.Cardinality())
	}
}
