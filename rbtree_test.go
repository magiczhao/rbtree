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
    ival = 7
    tree.root.left.right = newRbNode(ival)

    val, err = tree.Find(ival)
    if err != nil || val != ival {
        t.Error("Search failed")
    }
}

func TestRbTreeInsertEmpty(t *testing.T) {
    tree := RbTree{}
    var ival IntValue

    ival = 100
    tree.Insert(ival)
    if tree.root == nil {
        t.Error("Root is nil")
    }
    if tree.root.key != ival {
        t.Error("Root value is failed")
    }

    if tree.root.IsRed() {
        t.Error("Root Color is wrong")
    }
}

func TestRbTreeInsert(t *testing.T) {
    tree := RbTree{}
    var ival IntValue
    ival = 100

    tree.Insert(ival)
    ival = 10
    tree.Insert(ival)
    ival = 12
    tree.Insert(ival)
    ival = 33
    tree.Insert(ival)
    if tree.root.IsRed() {
        t.Error("Root Color is wrong")
    }
    _, err := tree.Find(ival)
    if err != nil {
        t.Error("Find failed")
    }

    ival = 199
    _, err = tree.Find(ival)
    if err == nil {
        t.Error("Find failed")
    }
    tree.Insert(ival)
    _, err = tree.Find(ival)
    if err != nil {
        t.Error("Find failed")
    }
    ival = 129
    tree.Insert(ival)
    if tree.root.IsRed() {
        t.Error("Root Color is wrong")
    }
}

func TestRbTreeMaxMin(t *testing.T) {
    tree := RbTree{}
    var ival IntValue
    ival = 12
    tree.Insert(ival)

    ival = 22
    tree.Insert(ival)

    ival = 100
    tree.Insert(ival)

    ival = 101
    tree.Insert(ival)

    ival = 102
    tree.Insert(ival)

    ival = 12
    if tree.Min() != ival {
        t.Error("Min failed")
    }

    ival = 102
    if tree.Max() != ival {
        t.Error("Max failed")
    }
}
