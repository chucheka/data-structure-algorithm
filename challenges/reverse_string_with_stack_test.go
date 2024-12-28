package challenges

import "testing"

func TestReverseString(t *testing.T) {
	want := "ekihC"

	got, err := ReverseString("Chike")

	if err != nil {
		t.Errorf("string reversal failed")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
