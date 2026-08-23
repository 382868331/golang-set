package mapset

import "testing"

func TestTask012FilterEven(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3, 4)
	f := s.Filter(func(v int) bool { return v%2 == 0 })
	if !f.Equal(NewThreadUnsafeSet(2, 4)) {
		t.Fatalf("filter=%v", f)
	}
}
