package mapset

import "testing"

func TestTask008ContainsOnePresent(t *testing.T) {
	s := NewThreadUnsafeSet("a", "b")
	if !s.ContainsOne("a") {
		t.Fatal("a reported absent")
	}
}
