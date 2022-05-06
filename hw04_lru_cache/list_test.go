package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push front", func(t *testing.T) {
		l := NewList()
		item1 := l.PushFront(10)
		item2 := l.PushFront(20)

		require.Equal(t, 2, l.Len())
		require.Equal(t, item2, l.Front())
		require.Equal(t, item1, l.Back())
	})

	t.Run("push back", func(t *testing.T) {
		l := NewList()
		item1 := l.PushBack(10)
		item2 := l.PushBack(20)

		require.Equal(t, 2, l.Len())
		require.Equal(t, item1, l.Front())
		require.Equal(t, item2, l.Back())
	})

	t.Run("push back and remove in list", func(t *testing.T) {
		l := NewList()
		item := l.PushBack(10)

		l.Remove(item)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push front and remove in list", func(t *testing.T) {
		l := NewList()
		item := l.PushFront(10)

		l.Remove(item)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("remove from empty list", func(t *testing.T) {
		l := NewList()

		l.Remove(nil)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("remove element not from list", func(t *testing.T) {
		l := NewList()
		item := &ListItem{Value: 10}

		l.Remove(item)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("move to front from empty list", func(t *testing.T) {
		l := NewList()

		l.MoveToFront(nil)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("move to front element not from list", func(t *testing.T) {
		l := NewList()
		item := &ListItem{Value: 10}

		l.MoveToFront(item)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()
		item1 := l.PushFront(10)
		item2 := l.PushFront(20)

		l.MoveToFront(item1)

		require.Equal(t, 2, l.Len())
		require.Equal(t, item1.Value, l.Front().Value)
		require.Equal(t, item2.Value, l.Back().Value)
	})

	t.Run("some operations move to front", func(t *testing.T) {
		n := 20
		l := NewList()
		l.PushBack(10)
		l.PushBack(20)

		for i := 0; i < n; i++ {
			l.MoveToFront(l.Back())
			l.MoveToFront(l.Back())
		}

		require.Equal(t, 2, l.Len())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
