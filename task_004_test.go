package mapset

import "testing"

func TestTask004EmptyCardinality(t *testing.T) {
	s := NewThreadUnsafeSet[int]()
	if s.Cardinality() != 0 {
		t.Fatalf("cardinality=%d", s.Cardinality())
	}
}
func TestTask004NonEmptyCardinality(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3)
	if s.Cardinality() != 3 {
		t.Fatalf("cardinality=%d", s.Cardinality())
	}
}
