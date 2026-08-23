package mapset

import "testing"

func TestTask003AppendFromDisjoint(t *testing.T) {
	s := NewThreadUnsafeSet(1)
	o := NewThreadUnsafeSet(2, 3)
	if n := s.AppendFrom(o); n != 2 || !s.Contains(1, 2, 3) {
		t.Fatalf("n=%d set=%v", n, s)
	}
}
