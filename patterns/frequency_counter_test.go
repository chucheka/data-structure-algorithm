package patterns

import "testing"

func TestSameSquared(t *testing.T) {

	var tests = []struct {
		name    string
		nums    []int
		squares []int
		want    bool
	}{
		{"[1, 2, 3], [4, 1, 9]", []int{1, 2, 3}, []int{4, 1, 9}, true},
		{"[1, 2, 3], [1, 9]", []int{1, 2, 3}, []int{1, 9}, false},
		{"[1, 2, 1], [4, 4, 1]", []int{1, 2, 1}, []int{4, 4, 1}, false},
		{"[2, 3, 6, 8, 8], [64, 36, 4, 9, 64]", []int{2, 3, 6, 8, 8}, []int{64, 36, 4, 9, 64}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SameSquared(tt.nums, tt.squares)
			if got != tt.want {
				t.Errorf("got %t want %t", got, tt.want)
			}
		})
	}

}

func TestValidAnagram(t *testing.T) {

	var tests = []struct {
		name string
		str1 string
		str2 string
		want bool
	}{
		{"silent and listen ----> true", "silent", "listen", true},
		{"martin and nitram ----> true", "martin", "nitram", true},
		{"cat and tag ----> false", "cat", "tag", false},
		{"rat and tar ----> true", "rat", "tar", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidAnagram(tt.str1, tt.str2)

			if got != tt.want {
				t.Errorf("got %t want %t", got, tt.want)
			}
		})
	}

}
