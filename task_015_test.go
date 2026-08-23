package mapset

import "testing"

func TestTask015EmptySet(t *testing.T) {
	if !NewThreadUnsafeSet[int]().IsEmpty() {
		t.Fatal("empty set reported non-empty")
	}
}
func TestTask015NonEmptySet(t *testing.T) {
	if NewThreadUnsafeSet(1).IsEmpty() {
		t.Fatal("non-empty set reported empty")
	}
}
