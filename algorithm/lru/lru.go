package lru

import (
	"container/list"
)

type Key interface{}
type Value interface{}

type Node struct {
	key   Key
	value Value
}

type Cache struct {
	queue    *list.List
	elements map[Key]*list.Element
	cap      int
}

func NewCache(cap int) *Cache {
	return &Cache{
		queue:    list.New(),
		elements: make(map[Key]*list.Element),
		cap:      cap,
	}
}

func (c *Cache) Add(key, value interface{}) {
	if c.elements == nil {
		c.elements = make(map[Key]*list.Element)
		c.queue = list.New()
	}
	if element, ok := c.elements[key]; ok {
		c.queue.MoveToFront(element)
		element.Value.(*Node).value = value
		return
	}

	element := c.queue.PushFront(&Node{key: key, value: value})
	c.elements[key] = element
	if c.cap != 0 && c.queue.Len() > c.cap {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	element, ok := c.elements[key]
	if ok {
		c.queue.MoveToFront(element)
		return element.Value.(*Node).value, ok
	}
	return nil, ok
}

func (c *Cache) RemoveOldest() {
	if c.queue == nil {
		return
	}
	element := c.queue.Back()
	c.queue.Remove(element)
	delete(c.elements, element.Value.(*Node).key)
}

func (c *Cache) Len() int {
	if c.queue == nil {
		return 0
	}
	return c.queue.Len()
}

func (c *Cache) Clear() {
	c.queue = nil
	c.elements = nil
}
