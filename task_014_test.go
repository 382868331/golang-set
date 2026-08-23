package mapset

import "testing"

func TestTask014IntersectSmallerLeft(t *testing.T) {
	a := NewThreadUnsafeSet(1, 2)
	b := NewThreadUnsafeSet(2, 3, 4)
	x := a.Intersect(b)
	if !x.Equal(NewThreadUnsafeSet(2)) {
		t.Fatalf("intersection=%v", x)
	}
}
