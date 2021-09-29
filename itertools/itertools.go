package itertools

import (
	"container/list"
)

type ListIter struct {
	items *list.List
	item  *list.Element
}

func New(items *list.List) *ListIter {
	iter := new(ListIter)
	iter.Init(items)
	return iter
}

func (l *ListIter) Init(items *list.List) {
	l.items = items
}

func (l *ListIter) Item() *list.Element {
	return l.item
}

func (l *ListIter) Next() bool {
	if l.item == nil {
		l.item = l.items.Front()
		return true
	}
	next := l.item.Next()
	if next != nil {
		l.item = next
		return true
	}
	return false
}
