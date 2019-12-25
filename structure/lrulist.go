package structure

type lruList struct {
	l   *list
	cap int
}

// NewLruList returns List which cap
func NewLruList(cap int) List {
	if cap == 0 {
		return nil
	}
	return &lruList{
		l:   &list{},
		cap: cap,
	}
}

func (lr *lruList) Insert(value interface{}) {
	if lr.l.size == lr.cap {
		lr.removeOldest()
	}
	lr.l.Insert(value)
}

func (lr *lruList) Remove(n Node) {
	lr.l.Remove(n)
}

func (lr *lruList) GetSize() int {
	return lr.l.size
}

func (lr *lruList) GetFirstNode() Node {
	return lr.l.firstNode
}

func (lr *lruList) GetLastNode() Node {
	return lr.l.lastNode
}

func (lr *lruList) removeOldest() {
	removeNode := lr.l.firstNode
	lr.l.Remove(removeNode)
}
