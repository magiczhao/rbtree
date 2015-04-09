package rbtree

import (
	"errors"
)

// RbColor is color for rbtree node
type RbColor bool

// IsRed test if the object is red colored
func (c RbColor) IsRed() bool {
	return c == true
}

// IsBlack test if the object is black colored
func (c RbColor) IsBlack() bool {
	return c == false
}

// SetBlack set the color of the object to black
func (c *RbColor) SetBlack() {
	*c = false
}

// SetRed set the color of the object to red
func (c *RbColor) SetRed() {
	*c = true
}

// Comparable is the interface to key of the tree
type Comparable interface {
	// Less returns true if this object is less than argument
	Less(Comparable) bool
}

// NInf is a predefined Comparable
type NInf struct{}

// Less is always true
func (n NInf) Less(data Comparable) bool {
	return true
}

// predefined objects
var (
	ninf               = NInf{}
	ErrorNotFound      = errors.New("value not found")
	ErrorAlreadyExists = errors.New("Node already exists")
)

// rbNode is node in tree
type rbNode struct {
	RbColor
	// tree struct
	parent *rbNode
	left   *rbNode
	right  *rbNode
	// key and value are stored together
	key Comparable
}

// RbTree is redblack tree
type RbTree struct {
	root *rbNode
}

// Max get the max valued object in rbtree
func (t RbTree) Max() Comparable {
	if t.root == nil {
		return nil
	}
	node := t.root
	for node.right != nil {
		node = node.right
	}
	return node.key
}

// Min get the min valued object in rbtree
func (t RbTree) Min() Comparable {
	if t.root == nil {
		return nil
	}
	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node.key
}

// newRbNode malloc and init a rbNode with data
func newRbNode(data Comparable) *rbNode {
	node := rbNode{false, nil, nil, nil, data}
	return &node
}

// Insert insert a new value into the tree
func (t *RbTree) Insert(data Comparable) error {
	// 1. root is empty
	if t.root == nil {
		// empty tree
		t.root = newRbNode(data)
		t.root.SetBlack()
		return nil
	}
	node := t.findPosition(data)
	if node == nil {
		return ErrorNotFound
	}
	// init inserted node
	nNode := newRbNode(data)
	nNode.parent = node
	nNode.SetRed()
	// insert node
	if node.key.Less(data) {
		node.right = nNode
	} else if data.Less(node.key) {
		node.left = nNode
	} else {
		// already in tree
		return ErrorAlreadyExists
	}
	if node.IsBlack() {
		return nil
	}
	for node != t.root && node.IsRed() {
		uncle := node.parent.right
		if node == node.parent.right {
			uncle = node.parent.left
		}
		// 2. uncle is red
		if uncle != nil && uncle.IsRed() {
			node.SetBlack()
			uncle.SetBlack()
			if node.parent != t.root {
				node.parent.SetRed()
			}
			nNode, node = node.parent, nNode.parent
			continue
		}
		// 3. uncle is black
		if nNode == node.right && node == node.parent.left {
			t.rotateLeft(node)
			node, nNode = nNode, node
		} else if nNode == node.right && node == node.parent.right {
			node.SetBlack()
			node.parent.SetRed()
			if node.parent == t.root {
				t.root = node
			}
			t.rotateLeft(node.parent)
			break
		} else if nNode == node.left && node == node.parent.left {
			node.SetBlack()
			node.parent.SetRed()
			if node.parent == t.root {
				t.root = node
			}
			t.rotateRight(node.parent)
			break
		} else { // nNode == node.left && node = node.parent.right
			t.rotateRight(node)
			node, nNode = nNode, node
		}
	}
	return nil
}

// Find search the entire tree to find the data, if not found error will
// set to ErrorNotFound
func (t RbTree) Find(data Comparable) (Comparable, error) {
	node := t.findPosition(data)
	if node == nil {
		return ninf, ErrorNotFound
	}
	// Equal
	if !data.Less(node.key) && !node.key.Less(data) {
		return node.key, nil
	}
	return ninf, ErrorNotFound
}

func (t *RbTree) rotateLeft(node *rbNode) {
	parent := node.parent
	if node.right == nil {
		return
	}
	right := node.right
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	node.parent = right
	right.left = node
	right.parent = parent
	if parent == nil {
		return
	}
	if node == parent.left {
		parent.left = right
	} else {
		parent.right = right
	}
}

func (t *RbTree) rotateRight(node *rbNode) {
	parent := node.parent
	if node.left == nil {
		return
	}
	left := node.left
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	node.parent = left
	left.right = node
	left.parent = parent
	if parent == nil {
		return
	}
	if node == parent.left {
		parent.left = left
	} else {
		parent.right = left
	}
}

func (t RbTree) findPosition(data Comparable) *rbNode {
	if t.root == nil {
		return nil
	}
	node := t.root
	parent := node
	for node != nil {
		parent = node
		if data.Less(node.key) {
			node = node.left
		} else if node.key.Less(data) {
			node = node.right
		} else {
			// node is the target
			return node
		}
	}
	return parent
}

func allChildIsBlack(node *rbNode) bool {
	if node == nil {
		return true
	}
	if node.left != nil && node.left.IsRed() {
		return false
	}
	if node.right != nil && node.right.IsRed() {
		return false
	}
	return true
}

func swapColor(node1 *rbNode, node2 *rbNode) {
	color1 := false
	if node1.IsRed() {
		color1 = true
	}
	color2 := false
	if node2.IsRed() {
		color2 = true
	}
	if color1 {
		node1.SetRed()
	} else {
		node1.SetBlack()
	}
	if color2 {
		node2.SetRed()
	} else {
		node2.SetBlack()
	}
}

// Delete removes the data from the tree
func (t *RbTree) Delete(data Comparable) {
	node := t.findPosition(data)
	if node == nil {
		return
	}
	if node.left != nil && node.right != nil {
		// find the max value in left child
		target := node.left
		for target.right != nil {
			target = target.right
		}
		// exchange value
		node.key, target.key = target.key, node.key
		node = target
	}
	// 1. remove the root
	if node == t.root {
		t.root = nil
		return
	}
	// remove the node
	parent := node.parent
	child := node.left
	if child == nil {
		child = node.right
	}

	if parent.left == node {
		parent.left = child
	} else {
		parent.right = child
	}
	if child != nil {
		child.parent = parent
	}
	// 2. node is red, just return
	if node.IsRed() {
		return
	}

	// Adjust tree
	if child != nil {
		node = child
	} else {
		node = parent
	}

	for node != t.root {
		parent = node.parent
		// 3. node is red
		if node.IsRed() {
			node.SetBlack()
			return
		}
		sibling := parent.left
		if node == parent.left {
			sibling = parent.right
		}
		// 4. no sibling
		if sibling == nil {
			node = parent
			continue
		}
		// 4. sibling is red
		if sibling != nil && sibling.IsRed() {
			parent.SetRed()
			sibling.SetBlack()
			if sibling == parent.left {
				t.rotateRight(parent)
			} else {
				t.rotateLeft(parent)
			}
			if parent == t.root {
				t.root = sibling
			}
			node = parent
			continue
		}
		// 5. sibling is black & all child is black
		if sibling != nil && sibling.IsBlack() && allChildIsBlack(sibling) {
			sibling.SetRed()
			node = parent
			continue
		}
		// 6. sibling is black & at least one child is red
		if sibling != nil && sibling.IsBlack() {
			if sibling == parent.right {
				// right-right case
				if sibling.right != nil && sibling.right.IsRed() {
					t.rotateLeft(parent)
					swapColor(parent, sibling)
					sibling.right.SetBlack()
					return
				} else if sibling.left != nil && sibling.left.IsRed() {
					// right-left case
					t.rotateRight(sibling)
					sibling.left.SetBlack()
					sibling.SetRed()
					continue
				}
			} else {
				// left-left case
				if sibling.left != nil && sibling.left.IsRed() {
					t.rotateRight(parent)
					swapColor(parent, sibling)
					sibling.right.SetBlack()
					return
				} else if sibling.right != nil && sibling.right.IsRed() {
					// left-right case
					t.rotateLeft(sibling)
					sibling.right.SetBlack()
					sibling.SetRed()
					continue
				}
			}
		}
	}
}
