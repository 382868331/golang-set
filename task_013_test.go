package mapset

import "testing"

func TestTask013EqualSameMembers(t *testing.T) {
	a := NewThreadUnsafeSet(1, 2)
	b := NewThreadUnsafeSet(2, 1)
	if !a.Equal(b) {
		t.Fatal("same members not equal")
	}
}
