package mapset

import "testing"

func TestTask002AppendCount(t *testing.T) {
	s := NewThreadUnsafeSet(1)
	n := s.Append(1, 2, 3)
	if n != 2 {
		t.Fatalf("added=%d, want 2", n)
	}
	if s.Cardinality() != 3 {
		t.Fatalf("cardinality=%d", s.Cardinality())
	}
}
