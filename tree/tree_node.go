package tree

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data: data}
}
