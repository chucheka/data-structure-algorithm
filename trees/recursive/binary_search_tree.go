package recursive

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func Contains(root *Node, data int) bool {

	if root == nil {
		return false
	}

	if root.data == data {
		return true
	}

	if data < root.data {
		return Contains(root.left, data)
	} else {
		return Contains(root.right, data)
	}
}

func (bst *BinarySearchTree) Insert(data int) {
	bst.insertRec(bst.root, data)
}

func (bst *BinarySearchTree) Delete(data int) {

	deleteRec(bst.root, data)
}

func deleteRec(temp *Node, data int) *Node {

	if temp == nil {
		return nil
	}

	if data < temp.data {
		temp.left = deleteRec(temp.left, data)
	} else if data > temp.data {
		temp.right = deleteRec(temp.right, data)
	} else {
		/*
			THERE ARE 4 SCENARIOS WHEN data == temp.data
			(1) temp node is a left node
			(2) temp node has only left child
			(3) temp node has only right child
			(4) temp has both left and right children
		*/

		if temp.left == nil && temp.right == nil {
			return nil
		} else if temp.left == nil {
			temp = temp.right
		} else if temp.right == nil {
			temp = temp.left
		} else {

			// the min node is always opened on the left
			subTreeMin := minValue(temp.right)
			temp.data = subTreeMin
			temp.right = deleteRec(temp.right, subTreeMin)
		}
	}
	return temp
}

func (bst *BinarySearchTree) insertRec(node *Node, data int) *Node {

	if bst.root == nil {
		bst.root = &Node{
			data: data,
		}
		return bst.root
	}

	if node == nil {
		return &Node{data, nil, nil}
	}

	if data <= node.data {
		node.left = bst.insertRec(node.left, data)
	}
	if data > node.data {
		node.right = bst.insertRec(node.right, data)
	}
	return node

}

func minValue(currentNode *Node) int {

	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode.data
}
