package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data    any
	Pointer *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[string]*Node
	Capacity int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		Queue: list.New(), 
		Items: make(map[string]*Node),
		Capacity: capacity,
	}
}

func (l *LRUCache) Set(key string, value any) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(string))
		}
		l.Items[key] = &Node{Data: value, Pointer: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.Pointer)
	}
}

func (l *LRUCache) Get(key string) any {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.Pointer)
		return item.Data
	}

	return -1
}

func main()  {
	fmt.Println("==> LRUCache Test Case <==")
	obj := NewLRUCache(2)
	obj.Set("joan", 26)
	obj.Set("patricia", 28)
	obj.Set("wilson", 25)
	obj.Set("fulano", 20)
	
	fmt.Println(obj.Get("joan"))
	fmt.Println(obj.Get("patricia"))
	fmt.Println(obj.Get("wilson"))
	fmt.Println(obj.Get("fulano"))
}