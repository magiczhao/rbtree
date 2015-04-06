package rbtree

import (
	"testing"
)

type IntValue int

func (i IntValue) Less(val Comparable) bool {
	return i < val.(IntValue)
}

func TestRbTreeFind(t *testing.T) {
	tree := RbTree{}
	var ival IntValue
	ival = 10
	tree.root = newRbNode(ival)
	ival = 6
	tree.root.left = newRbNode(ival)
	val, err := tree.Find(ival)
	if err != nil {
		t.Error("Failed to find", ival)
	}
	if val != ival {
		t.Error("Wrong value")
	}
	ival = 5
	_, err = tree.Find(ival)
	if err == nil {
		t.Error("Should not find item")
	}
}
