package calc

import "testing"

func TestAdd(t *testing.T) {
	want := 12
	got := Add(5, 7)

	if got != want {
		t.Errorf("expected %q but add returned %q", want, got)
	}
}
