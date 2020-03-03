package structure


type lruList struct {
	l   *List
	cap int
}

// newLruList returns List which size
func newLruList(cap int) *lruList {
	if cap == 0 {
		return nil
	}
	return &lruList{
		l:   NewList(),
		cap: cap,
	}
}

func (lr *lruList) Insert(value interface{}) {
	if lr.l.Size == lr.cap {
		lr.removeOldest()
	}
	lr.l.Insert(value)
}

func (lr *lruList) Remove(n *Node) {
	lr.l.Remove(n)
}

func (lr *lruList) GetSize() int {
	return lr.l.Size
}

func (lr *lruList) GetFirstNode() *Node {
	return lr.l.FirstNode
}

func (lr *lruList) GetLastNode() *Node {
	return lr.l.LastNode
}

func (lr *lruList) removeOldest() {
	removeNode := lr.l.FirstNode
	lr.l.Remove(removeNode)
}
