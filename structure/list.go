package structure

/*
  singly and doubly List
*/

type SinglyList struct {
	FirstNode *SinglyNode
	LastNode  *SinglyNode
	Size      int
}

type SinglyNode struct {
	NextNode *SinglyNode
	Value    interface{}
}

// NewSingleList return SinglyList interface.
func NewSingleList() *SinglyList {
	return new(SinglyList)
}

// Insert add new Node with giving Value.
func (l *SinglyList) Insert(value interface{}) {
	newNode := &SinglyNode{
		Value: value,
	}
	if l.Size != 0 {
		l.LastNode.NextNode = newNode
	} else {
		l.FirstNode = newNode
	}
	l.LastNode = newNode
	l.Size++
}

func (l *SinglyList) Remove(n *SinglyNode) {
	if n == nil {
		return
	}
	node := l.FirstNode
	var pre *SinglyNode
	for ; ; node = node.NextNode {
		if node == nil {
			return
		}
		if node == n {
			if pre == nil {
				l.FirstNode = l.FirstNode.NextNode
			} else {
				pre.NextNode = node.NextNode
			}
			break
		}

		pre = node
	}
	l.Size--
}

// List is doubly List
type List struct {
	FirstNode *Node
	LastNode  *Node
	Size      int
}

// Node is doubly List Node
type Node struct {
	NextNode *Node
	PreNode  *Node
	Value    interface{}
}

// NewList return doubly List
func NewList() *List {
	return new(List)
}

func (l *List) Insert(value interface{}) {
	newNode := &Node{
		Value: value,
	}
	if l.Size != 0 {
		newNode.PreNode = l.LastNode
		l.LastNode.NextNode = newNode
	} else {
		l.FirstNode = newNode
	}
	l.LastNode = newNode
	l.Size++
}

func (l *List) Remove(n *Node) {
	if n == nil {
		return
	}
	defer func() {
		l.Size--
	}()
	// Remove first Node
	if n.PreNode == nil {
		l.FirstNode = n.NextNode
		if l.FirstNode != nil {
			n.NextNode.PreNode = nil
		}
		return
	}

	// Remove last Node
	if n.NextNode == nil {
		l.LastNode = n.PreNode
		if l.LastNode != nil {
			n.PreNode.NextNode = nil
		}
		return
	}

	// Node is middle in List
	n.PreNode.NextNode = n.NextNode
	n.NextNode.PreNode = n.PreNode
}
