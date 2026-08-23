package mapset

import "testing"

func TestTask015EmptySet(t *testing.T) {
	if !NewThreadUnsafeSet[int]().IsEmpty() {
		t.Fatal("empty set reported non-empty")
	}
}
