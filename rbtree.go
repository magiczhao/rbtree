package rbtree

import (
    "fmt"
	"errors"
)

type RbColor bool

func (c RbColor) IsRed() bool {
	return c == true
}

func (c RbColor) IsBlack() bool {
	return c == false
}

func (c *RbColor) SetBlack() {
	*c = false
}

func (c *RbColor) SetRed() {
	*c = true
}

type Comparable interface {
	Less(Comparable) bool
}

type NInf struct{}

func (n NInf) Less(data Comparable) bool {
	return true
}

var (
	ninf               = NInf{}
	ErrorNotFound      = errors.New("value not found")
	ErrorAlreadyExists = errors.New("Node already exists")
)

type rbNode struct {
	RbColor
	// tree struct
	parent *rbNode
	left   *rbNode
	right  *rbNode
	// key and value are stored together
	key Comparable
}

type RbTree struct {
	root *rbNode
}

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

func newRbNode(data Comparable) *rbNode {
	node := rbNode{false, nil, nil, nil, data}
	return &node
}

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
			nNode = node.parent
			node = nNode.parent
			continue
		}
		// 3. uncle is black
		if nNode == node.right && node == node.parent.left {
            fmt.Println("Befaore", nNode, node, node.parent)
			t.rotateLeft(node)
            fmt.Println("After", nNode, node, node.parent)
			node, nNode = nNode, node
		} else if nNode == node.right && node == node.parent.right {
            t.rotateLeft(node.parent)
            node.SetBlack()
            node.parent.SetRed()
            if node.parent == t.root {
                t.root = node
            }
		    break
        } else if nNode == node.left && node == node.parent.left {
            fmt.Println(nNode, node, node.parent)
            t.rotateRight(node.parent)
            node.SetBlack()
            node.parent.SetRed()
            if node.parent == t.root {
                t.root = node
            }
            break
        } else { // nNode == node.left && node = node.parent.right
            t.rotateRight(node)
            node, nNode = nNode, node
        }
	}
	return nil
}

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
