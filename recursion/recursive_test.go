package recursion

import "testing"

func TestFactorial(t *testing.T) {

	want := uint(24)

	got := Factorial(4)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestPow(t *testing.T) {

	testCases := []struct {
		name  string
		num   uint
		index uint
		want  uint
	}{{"num to power of 0", 3, 0, 1}, {"num to power of 1", 3, 1, 3}, {"num to power of 3", 3, 3, 27}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Pow(tt.num, tt.index)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
