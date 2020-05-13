package tree

type node struct {
	left  *node
	right *node
	data  int
}

type Tree struct {
	Root *node
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &node{data: value}
		return
	}

	insert(t.Root, value)
}

func insert(n *node, value int) {
	if n.data > value {
		if n.left == nil {
			n.left = &node{data: value}
			return
		}

		insert(n.left, value)
	} else {
		if n.right == nil {
			n.right = &node{data: value}
			return
		}

		insert(n.right, value)
	}
}
