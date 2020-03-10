package btree

type (
	btree struct {
		value int
		left  *btree
		right *btree
	}
)

func newBST(root int, values []int) *btree {
	tree := &btree{value: root}
	for _, value := range values {
		tree = tree.insertRecursive(value)
	}
	return tree
}

func (node *btree) insertRecursive(value int) *btree {
	if value < node.value {
		if node.left != nil {
			node.left.insertRecursive(value)
		} else {
			node.left = &btree{value: value}
		}
	} else {
		if node.right != nil {
			node.right.insertRecursive(value)
		} else {
			node.right = &btree{value: value}
		}
	}
	return node
}

func (node *btree) deleteRecursive(value int) *btree {
	node.removeRecursive(value, nil)
	return node
}

func (node *btree) removeRecursive(value int, parent *btree) {
	// Find node with value to be deleted
	if value < node.value {
		if node.left != nil {
			node.left.removeRecursive(value, node)
		}
	} else if value > node.value {
		if node.right != nil {
			node.right.removeRecursive(value, node)
		}
		// If value not greater or less we have an hit
	} else {
		// case 1 if node has both left and rigt node.
		if node.left != nil && node.right != nil {
			node.value = node.right.minValueRecursive()
			node.right.removeRecursive(node.value, node)
		} else if parent == nil {
			// case 2 node that we want to removce is root node
			if node.left != nil {
				node.value = node.left.value
				node.right = node.left.right
				node.left = node.left.left
			} else if node.right != nil {
				node.value = node.right.value
				node.left = node.right.left
				node.right = node.right.right
			}
		} else if parent.left == node {
			if node.left != nil {
				parent.left = node.left
			} else {
				parent.left = node.right
			}

		} else if parent.right == node {
			if node.right != nil {
				parent.right = node.right
			} else {
				parent.right = node.left
			}
		}
	}
}

func (node *btree) minValueRecursive() int {
	if node.left == nil {
		return node.value
	}
	return node.left.minValueRecursive()
}

func (node *btree) maxValueRecursive() int {
	if node.right == nil {
		return node.value
	}
	return node.right.maxValueRecursive()
}

func (node *btree) inOrderTraverseRecursive(array []int) []int {
	if node.left != nil {
		array = node.left.inOrderTraverseRecursive(array)
	}
	array = append(array, node.value)
	if node.right != nil {
		array = node.right.inOrderTraverseRecursive(array)
	}
	return array
}
