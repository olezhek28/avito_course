package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	var found bool
	var item *ListItem
	if item, found = l.items[key]; found {
		item.Value = value
		l.queue.MoveToFront(item)

		return found
	}

	if l.capacity == l.queue.Len() {
		last := l.queue.Back()
		//delete(l.items, last.Value)
		l.queue.Remove(last)
	}

	l.items[key] = l.queue.PushFront(value)
	return found
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	item := l.items[key]
	if item == nil {
		return nil, false
	}

	l.queue.MoveToFront(item)
	return item.Value, true
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}
