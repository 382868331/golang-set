package mapset

import "testing"

func TestTask011EachVisitsUntilStop(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3)
	n := 0
	s.Each(func(int) bool { n++; return false })
	if n != 3 {
		t.Fatalf("visits=%d", n)
	}
}
func TestTask011EachStopsOnTrue(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3)
	n := 0
	s.Each(func(int) bool { n++; return true })
	if n != 1 {
		t.Fatalf("visits=%d", n)
	}
}
