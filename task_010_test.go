package mapset

import "testing"

func TestTask010DifferenceOverlap(t *testing.T) {
	a := NewThreadUnsafeSet(1, 2, 3)
	b := NewThreadUnsafeSet(2, 4)
	d := a.Difference(b)
	if !d.Equal(NewThreadUnsafeSet(1, 3)) {
		t.Fatalf("difference=%v", d)
	}
}
func TestTask010DifferenceDisjoint(t *testing.T) {
	a := NewThreadUnsafeSet(1, 2)
	b := NewThreadUnsafeSet(8)
	d := a.Difference(b)
	if !d.Equal(a) {
		t.Fatalf("difference=%v", d)
	}
}
