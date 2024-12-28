package trees

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {

	bst := &BinarySearchTree{}

	bst.Insert(60)
	bst.Insert(40)
	bst.Insert(20)
	bst.Insert(90)
	bst.Insert(30)
	bst.Insert(50)
	bst.Insert(10)
	bst.Insert(80)
	bst.Insert(70)

	want := 20

	got := bst.root.left.left.data

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	if !bst.Contain(80) {
		t.Errorf("expect binary tree to contain 80")
	}

}
