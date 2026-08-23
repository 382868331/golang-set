package mapset

import "testing"

func TestTask009AnyOnePresent(t *testing.T) {
	s := NewThreadUnsafeSet(2, 4)
	if !s.ContainsAny(4) {
		t.Fatal("present candidate should match")
	}
}
