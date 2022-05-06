package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size     int
	first    *ListItem
	last     *ListItem
	valueMap map[interface{}]*ListItem
}

func NewList() List {
	return &list{
		valueMap: make(map[interface{}]*ListItem),
	}
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.first == nil {
		l.first = &ListItem{
			Value: v,
		}
		l.last = l.first
	} else {
		tmp := l.first
		l.first = &ListItem{
			Value: v,
			Next:  tmp,
		}

		tmp.Prev = l.first
	}

	l.size++
	l.valueMap[l.first.Value] = l.first

	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.last == nil {
		l.last = &ListItem{
			Value: v,
		}
		l.first = l.last
	} else {
		tmp := l.last
		l.last = &ListItem{
			Value: v,
			Prev:  tmp,
		}

		tmp.Next = l.last
	}

	l.size++
	l.valueMap[l.last.Value] = l.last

	return l.last
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	if ptr, found := l.valueMap[i.Value]; found {
		prevElem := ptr.Prev
		nexElem := ptr.Next

		if prevElem != nil {
			prevElem.Next = nexElem
		}
		if nexElem != nil {
			nexElem.Prev = prevElem
		}

		delete(l.valueMap, i.Value)
		l.size--

		if l.size == 0 {
			l.first = nil
			l.last = nil
		}
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}

	if ptr, found := l.valueMap[i.Value]; found {
		if ptr == l.first {
			return
		}

		prevElem := ptr.Prev
		nexElem := ptr.Next

		if prevElem != nil {
			prevElem.Next = nexElem
		}
		if nexElem != nil {
			nexElem.Prev = prevElem
		}

		l.PushFront(i.Value)
	}
}
