package sorting

import "testing"

func TestSelectionSort(t *testing.T) {

	nums := []int{5, 3, 6, 2, 10}

	want := 2

	got := SelectionSort(nums)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
