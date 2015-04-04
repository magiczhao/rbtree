package rbtree

type RbColor bool

func (c RbColor) IsRed() {
    return c
}

func (c RbColor) IsBlack() {
    return !c
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

type NInf struct {}

func (n NInf) Less(data Comparable) {
    return true
}

var (
    ninf = NInf{}
    ErrorNotFound = errors.New("value not found")
)

type rbNode struct {
    // tree struct
    parent *rbNode
    left *rbNode
    right *rbNode
    // color of the node
    color RbColor
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
    node := rbNode{nil, nil, nil, true, data}
    return &node
}

func (t *RbTree) Insert(data Comparable) error {
    node := t.findPosition(data)
    if node == nil {
        // empty tree
        t.root = newRbNode(data)
        return nil
    }
    nNode := newRbNode(data)
    nNode.color.SetRed()
    if node.key.Less(data) {
        node.right = nNode
    } else {
        node.left = nNode
    }
    if node.color.IsBlack() {
        return nil
    }
    return nil
}

func (t RbTree) Find(data Comparable) (Comparable, error) {
    node := t.findPosition(data)
    if node == nil {
        return ninf, ErrorNotFound
    }
    // Equal
    if !data.Less(node) && !node.key.Less(data) {
        return node.key
    }
    return ninf, ErrorNotFound
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
