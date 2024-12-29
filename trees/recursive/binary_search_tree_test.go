package recursive

import (
	"testing"
)

func TestBinarySearchTree_Contains(t *testing.T) {

	bst := &BinarySearchTree{}

	bst.Insert(20)
	bst.Insert(10)
	bst.Insert(40)

	/*
				20
		       /  \
		     10    40
	*/

	if !Contains(bst.root, 40) {
		t.Errorf("got false want true")
	}

}

func TestBinarySearchTree_Delete(t *testing.T) {

	bst := &BinarySearchTree{}

	bst.Insert(20)
	bst.Insert(10)
	bst.Insert(40)
	bst.Insert(5)
	bst.Insert(12)
	bst.Insert(60)

	/*
					20
			       /  \
			     10    40
				/  \    \
		       5    12    60
	*/

	bst.Delete(20)

	if Contains(bst.root, 20) {
		t.Errorf("expect BST not to still contain a deleted node")
	}

	want := 40

	got := bst.root.data

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
