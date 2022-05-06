package hw04lrucache

import "sync"

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
	mutex    *sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    &sync.Mutex{},
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	item := &cacheItem{
		key:   key,
		value: value,
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if listEl, found := l.items[key]; found {
		listEl.Value = item
		l.queue.MoveToFront(listEl)
		l.items[key] = l.queue.Front()

		return true
	}

	if l.capacity == l.queue.Len() {
		back := l.queue.Back()
		delete(l.items, back.Value.(*cacheItem).key)
		l.queue.Remove(back)
	}

	l.items[key] = l.queue.PushFront(item)
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	item := l.items[key]
	if item == nil {
		return nil, false
	}

	l.queue.MoveToFront(item)
	l.items[key] = l.queue.Front()

	return item.Value.(*cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}
