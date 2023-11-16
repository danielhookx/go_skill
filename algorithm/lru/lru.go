package lru

import (
	xlist "github.com/danielhookx/xcontainer/list"
)

type Node[K comparable, V any] struct {
	key   K
	value V
}

type Cache[K comparable, V any] struct {
	l        *xlist.List[Node[K, V]]
	elements map[K]*xlist.Element[Node[K, V]]
	cap      int
}

func NewCache[K comparable, V any](cap int) *Cache[K, V] {
	return &Cache[K, V]{
		l:        xlist.New[Node[K, V]](),
		elements: make(map[K]*xlist.Element[Node[K, V]]),
		cap:      cap,
	}
}

func (c *Cache[K, V]) Add(key K, value V) {
	if c.elements == nil {
		c.elements = make(map[K]*xlist.Element[Node[K, V]])
		c.l = xlist.New[Node[K, V]]()
	}
	if element, ok := c.elements[key]; ok {
		c.l.MoveToFront(element)
		element.Value.key = key
		element.Value.value = value
		return
	}

	element := c.l.PushFront(Node[K, V]{key: key, value: value})
	c.elements[key] = element
	if c.cap != 0 && c.l.Len() > c.cap {
		c.RemoveOldest()
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	element, ok := c.elements[key]
	if ok {
		c.l.MoveToFront(element)
		return element.Value.value, ok
	}
	return *new(V), ok
}

func (c *Cache[K, V]) RemoveOldest() {
	if c.l == nil {
		return
	}
	element := c.l.Back()
	c.l.Remove(element)
	delete(c.elements, element.Value.key)
}

func (c *Cache[K, V]) Len() int {
	if c.l == nil {
		return 0
	}
	return c.l.Len()
}

func (c *Cache[K, V]) Clear() {
	c.l = nil
	c.elements = nil
}
