package challenges

import "testing"

func TestBalancedParentheses(t *testing.T) {

	var tests = []struct {
		name string
		s    string
		want bool
	}{
		{"Balanced", "({})[]", true},
		{"Not Balanced", "{(})[]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BalancedParentheses(tt.s)
			if got != tt.want {
				t.Errorf("got %t want %t", got, tt.want)
			}
		})
	}

}
