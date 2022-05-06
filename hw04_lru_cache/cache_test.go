package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		val, ok := c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("bbb")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("multi set", func(t *testing.T) {
		c := NewCache(10)

		ok := c.Set("aaa", 10)
		require.False(t, ok)

		ok = c.Set("bbb", 20)
		require.False(t, ok)

		ok = c.Set("ccc", 30)
		require.False(t, ok)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, val, 10)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, val, 20)

		val, ok = c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, val, 30)
	})

	t.Run("over write", func(t *testing.T) {
		c := NewCache(10)

		ok := c.Set("aaa", 10)
		require.False(t, ok)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, val, 10)

		ok = c.Set("aaa", 20)
		require.True(t, ok)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, val, 20)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(2)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		wasInCache = c.Set("ccc", 300)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		val, ok = c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, 300, val)
	})

	t.Run("purge logic with frequency of use", func(t *testing.T) {
		n := 10
		c := NewCache(2)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		for i := 0; i < n; i++ {
			val, ok := c.Get("aaa")
			require.True(t, ok)
			require.Equal(t, 100, val)

			val, ok = c.Get("bbb")
			require.True(t, ok)
			require.Equal(t, 200, val)
		}

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		wasInCache = c.Set("ccc", 300)
		require.False(t, wasInCache)

		val, ok = c.Get("bbb")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, 300, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
