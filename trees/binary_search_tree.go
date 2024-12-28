package trees

type Node struct {
	left  *Node
	right *Node
	data  int
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(data int) {

	newNode := &Node{
		data: data,
	}

	if bst.root == nil {
		bst.root = newNode
		return
	}

	temp := bst.root

	for {
		if temp.data == newNode.data {
			return
		}

		if newNode.data < temp.data {
			if temp.left == nil {
				temp.left = newNode
				return
			}
			temp = temp.left
		} else {
			if temp.right == nil {
				temp.right = newNode
				return
			}
			temp = temp.right
		}

	}

}

func (bst *BinarySearchTree) Contain(data int) bool {

	if bst.root == nil {
		return false
	}

	temp := bst.root

	for {
		if temp.data == data {
			return true
		}
		if data < temp.data {
			if temp.left == nil {
				return false
			}
			temp = temp.left
		} else {
			if temp.right == nil {
				return false
			}
			temp = temp.right
		}
	}

}
