package mapset

import "testing"

func TestTask006CloneIndependentAdd(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2)
	c := s.Clone()
	c.Add(3)
	if s.Contains(3) {
		t.Fatal("clone Add changed source")
	}
}
func TestTask006CloneIndependentRemove(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2)
	c := s.Clone()
	c.Remove(1)
	if !s.Contains(1) {
		t.Fatal("clone Remove changed source")
	}
}
