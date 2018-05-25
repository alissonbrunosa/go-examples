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

func (this *Tree) Insert(value int) {
	if this.Root == nil {
		this.Root = &node{data: value}
		return
	}
	insert(this.Root, value)
}

func insert(root *node, value int) {
	if root.data > value {
		if root.left == nil {
			root.left = &node{data: value}
			return
		}
		insert(root.left, value)
	} else {
		if root.right == nil {
			root.right = &node{data: value}
			return
		}
		insert(root.right, value)
	}
}
