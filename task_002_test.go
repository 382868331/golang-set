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
func TestTask002AppendAllDuplicates(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2)
	if n := s.Append(2, 1); n != 0 {
		t.Fatalf("added=%d", n)
	}
}
