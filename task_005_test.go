package mapset

import "testing"

func TestTask005ClearMany(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3, 4)
	s.Clear()
	if !s.IsEmpty() {
		t.Fatalf("remaining=%v", s)
	}
}
func TestTask005ClearSingleton(t *testing.T) {
	s := NewThreadUnsafeSet(9)
	s.Clear()
	if !s.IsEmpty() {
		t.Fatalf("remaining=%v", s)
	}
}
