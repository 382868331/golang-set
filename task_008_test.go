package mapset

import "testing"

func TestTask008ContainsOnePresent(t *testing.T) {
	s := NewThreadUnsafeSet("a", "b")
	if !s.ContainsOne("a") {
		t.Fatal("a reported absent")
	}
}
func TestTask008ContainsOneMissing(t *testing.T) {
	s := NewThreadUnsafeSet("a")
	if s.ContainsOne("z") {
		t.Fatal("z reported present")
	}
}
