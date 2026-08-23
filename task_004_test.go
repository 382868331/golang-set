package mapset

import "testing"

func TestTask004EmptyCardinality(t *testing.T) {
	s := NewThreadUnsafeSet[int]()
	if s.Cardinality() != 0 {
		t.Fatalf("cardinality=%d", s.Cardinality())
	}
}
