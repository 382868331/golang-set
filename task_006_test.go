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
