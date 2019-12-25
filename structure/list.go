// singly and doubly list
package structure

type SinglyList interface {
	Insert(value interface{})
	Remove(node SinglyNode)
	GetSize() int
	GetFirstNode() SinglyNode
	GetLastNode() SinglyNode
}

type SinglyNode interface {
	GetValue() interface{}
	GetNextNode() SinglyNode
	IsNil() bool
}

type singlyList struct {
	firstNode *singlyNode
	lastNode  *singlyNode
	size      int
}

type singlyNode struct {
	nextNode *singlyNode
	value    interface{}
}

func (n *singlyNode) GetValue() interface{} {
	return n.value
}

func (n *singlyNode) GetNextNode() SinglyNode {
	return n.nextNode
}

func (n *singlyNode) IsNil() bool {
	return n == nil
}

func NewSingleList() SinglyList {
	return new(singlyList)
}

func (l *singlyList) Insert(value interface{}) {
	newNode := &singlyNode{
		value: value,
	}
	if l.size != 0 {
		l.lastNode.nextNode = newNode
	} else {
		l.firstNode = newNode
	}
	l.lastNode = newNode
	l.size++
}

func (l *singlyList) GetFirstNode() SinglyNode {
	return l.firstNode
}

func (l *singlyList) GetLastNode() SinglyNode {
	return l.lastNode
}

func (l *singlyList) Remove(n SinglyNode) {
	if n == nil {
		return
	}
	node := l.firstNode
	var pre *singlyNode
	for ; ; node = node.nextNode {
		if node == nil {
			return
		}
		if node == n.(*singlyNode) {
			if pre == nil {
				l.firstNode = l.firstNode.nextNode
			} else {
				pre.nextNode = node.nextNode
			}
			break
		}
		pre = node
	}
	l.size--
}

func (l *singlyList) GetSize() int {
	return l.size
}

type List interface {
	Insert(value interface{})
	Remove(node Node)
	GetSize() int
	GetFirstNode() Node
	GetLastNode() Node
}

type Node interface {
	GetValue() interface{}
	GetPreNode() Node
	GetNextNode() Node
	IsNil() bool
}

type list struct {
	firstNode *node
	lastNode  *node
	size      int
}

type node struct {
	nextNode *node
	preNode  *node
	value    interface{}
}

func (n *node) GetValue() interface{} {
	return n.value
}

func (n *node) GetPreNode() Node {
	return n.preNode
}

func (n *node) GetNextNode() Node {
	return n.nextNode
}

func (n *node) IsNil() bool {
	return n == nil
}

func NewList() List {
	return new(list)
}

func (l *list) Insert(value interface{}) {
	newNode := &node{
		value: value,
	}
	if l.size != 0 {
		newNode.preNode = l.lastNode
		l.lastNode.nextNode = newNode
	} else {
		l.firstNode = newNode
	}
	l.lastNode = newNode
	l.size++
}

func (l *list) GetFirstNode() Node {
	return l.firstNode
}

func (l *list) GetLastNode() Node {
	return l.lastNode
}

func (l *list) Remove(n Node) {
	no, ok := n.(*node)
	if !ok {
		return
	}
	defer func() {
		l.size--
	}()
	// Remove first node
	if no.preNode == nil {
		l.firstNode = no.nextNode
		if l.firstNode != nil {
			no.nextNode.preNode = nil
		}
		return
	}

	// Remove last node
	if no.nextNode == nil {
		l.lastNode = no.preNode
		if l.lastNode != nil {
			no.preNode.nextNode = nil
		}
		return
	}

	// node is middle in list
	no.preNode.nextNode = no.nextNode
	no.nextNode.preNode = no.preNode
}

func (l *list) GetSize() int {
	return l.size
}