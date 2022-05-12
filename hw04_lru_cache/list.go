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
	l.front = l.pushBefore(l.front, v)

	if l.back == nil {
		l.back = l.front
	}

	return l.front
}

func (l *list) pushBefore(node *ListItem, v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  node,
	}

	if node != nil {
		newItem.Prev = node.Prev
		node.Prev = newItem
	}

	l.size++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.back = l.pushAfter(l.back, v)

	if l.front == nil {
		l.front = l.back
	}

	return l.back
}

func (l *list) pushAfter(node *ListItem, v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Prev:  l.back,
	}

	if node != nil {
		newItem.Next = node.Next
		node.Next = newItem
	}

	l.size++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i == nil || l.Len() == 0 {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i == l.front {
		l.front = i.Next
	}
	if i == l.back {
		l.back = i.Prev
	}

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Len() == 0 || i == nil || i == l.front || (i.Prev == nil && i.Next == nil) {
		return
	}

	l.Remove(i)
	l.PushFront(i.Value)
}
