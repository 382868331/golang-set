package mapset

import "testing"

func TestTask007ContainsPresent(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3)
	if !s.Contains(1, 3) {
		t.Fatal("present elements reported absent")
	}
}
func TestTask007ContainsMissing(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2)
	if s.Contains(1, 9) {
		t.Fatal("missing element accepted")
	}
}
