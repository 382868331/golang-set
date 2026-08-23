package mapset

import "testing"

func TestTask018PopShrinks(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2)
	_, ok := s.Pop()
	if !ok {
		t.Fatal("Pop failed")
	}
	if s.Cardinality() != 1 {
		t.Fatalf("cardinality=%d", s.Cardinality())
	}
}
