package strings

import "testing"

func TestIsIsomorphicStrings(t *testing.T) {

	var tests = []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"egg and add", "egg", "add", true},
		{"foo and bar", "foo", "bar", false},
		{"paper and title", "paper", "title", true},
		{"fo and bar", "fo", "bar", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIsomorphicStrings(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %t want %t", got, tt.want)
			}
		})
	}

}
