package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data    int
	Pointer *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		Queue: list.New(), 
		Items: make(map[int]*Node),
		Capacity: capacity,
	}
}

func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}
		l.Items[key] = &Node{Data: value, Pointer: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.Pointer)
	}
}

func (l *LRUCache) Get(key int) int {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.Pointer)
		return item.Data
	}

	return -1
}

func main()  {
	fmt.Println("==> LRUCache Test Case <==")
	obj := NewLRUCache(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(4, 1)
	obj.Put(5, 7)
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(5))
}