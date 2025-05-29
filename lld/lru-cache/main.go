package main

import "fmt"

type Node struct {
	Key        int
	Value      int
	Prev, Next *Node
}

type LRUCache struct {
	Capacity   int
	Cache      map[int]*Node
	Head, Tail *Node
}

func NewNode(key int, value int) (node *Node) {
	node = &Node{
		Key:   key,
		Value: value,
	}
	return
}

func NewLRUCache(capacity int) (lruCache *LRUCache) {
	if capacity <= 0 {
		capacity = 3
	}
	lruCache = &LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*Node, capacity),
		Head:     &Node{},
		Tail:     &Node{},
	}
	lruCache.Head.Next = lruCache.Tail
	lruCache.Tail.Prev = lruCache.Head
	return
}

func (cache *LRUCache) Get(key int) (res int) {
	res = -1
	node, ok := cache.Cache[key]
	if ok {
		cache.moveToFront(node)
		res = node.Value
		return
	}
	return
}

func (cache *LRUCache) Put(key int, value int) (ok bool) {
	node, ok := cache.Cache[key]
	if ok {
		node.Value = value
		cache.moveToFront(node)
	}
	if len(cache.Cache) >= cache.Capacity {
		cache.removeLeastFrequentlyUsed()
	}
	newNode := NewNode(key, value)
	cache.Cache[key] = newNode
	cache.insertAtFront(newNode)
	return
}

func (cache *LRUCache) remove(node *Node) (ok bool) {
	// Deleting position can be anywhere
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	return
}

func (cache *LRUCache) insertAtFront(node *Node) (ok bool) {
	node.Prev = cache.Head
	node.Next = cache.Head.Next
	cache.Head.Next.Prev = node
	cache.Head.Next = node
	return
}

func (cache *LRUCache) moveToFront(node *Node) {
	cache.remove(node)
	cache.insertAtFront(node)
}

func (cache *LRUCache) removeLeastFrequentlyUsed() {
	lru := cache.Tail.Prev
	if lru != cache.Head {
		cache.remove(lru)
		delete(cache.Cache, lru.Key)
	}
}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
