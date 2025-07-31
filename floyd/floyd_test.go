package floyd

import (
	"testing"
)

func TestFloyd(t *testing.T) {
	want := "1\n2 3\n4 5 6\n7 8 9 10"
	got := Floyd(4)

	if got != want {
		t.Errorf("got\n%s\nwant\n%s", got, want)
	}
}
