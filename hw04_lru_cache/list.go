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
	front *ListItem
	back  *ListItem
	size  int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  l.front,
	}

	if l.front != nil {
		l.front.Prev = newItem
	}

	l.front = newItem
	l.size++

	if l.back == nil {
		l.back = l.front
	}

	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Prev:  l.back,
	}

	if l.back != nil {
		l.back.Next = newItem
	}

	l.back = newItem
	l.size++

	if l.front == nil {
		l.front = l.back
	}

	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	l.size--

	switch {
	case l.Len() <= 0:
		l.front = nil
		l.back = nil
		l.size = 0

	case i == l.front:
		l.front = l.front.Next

	case i == l.back:
		l.back = l.back.Prev

	default:
		prevElem := i.Prev
		nextElem := i.Next

		prevElem.Next = nextElem
		nextElem.Prev = prevElem
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || i == l.front || (i.Prev == nil && i.Next == nil) {
		return
	}

	prevElem := i.Prev
	nexElem := i.Next

	if prevElem != nil {
		prevElem.Next = nexElem
	}
	if nexElem != nil {
		nexElem.Prev = prevElem
	}

	l.PushFront(i.Value)
}
