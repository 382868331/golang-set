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
