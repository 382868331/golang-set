package mapset

import "testing"

func TestTask001AddNewReturnsTrue(t *testing.T) {
	s := NewThreadUnsafeSet[int]()
	if !s.Add(7) {
		t.Fatal("new Add returned false")
	}
	if !s.Contains(7) {
		t.Fatal("item missing")
	}
}
func TestTask001AddDuplicateReturnsFalse(t *testing.T) {
	s := NewThreadUnsafeSet(7)
	if s.Add(7) {
		t.Fatal("duplicate Add returned true")
	}
	if s.Cardinality() != 1 {
		t.Fatalf("cardinality=%d", s.Cardinality())
	}
}
