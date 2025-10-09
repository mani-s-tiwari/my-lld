package main

import "fmt"

// Node represents each element in the doubly linked list
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache implements the cache logic using a doubly linked list + hashmap
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

// Constructor initializes the cache
func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// remove disconnects a node from the linked list
func (lrucaching *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// add inserts a node right after head (most recently used)
func (lrucaching *LRUCache) add(node *Node) {
	prev, next := lrucaching.head, lrucaching.head.next
	prev.next = node
	node.prev = prev
	node.next = next
	next.prev = node
}


// moveToHead moves an existing node to the most recently used position
func (lrucaching *LRUCache) moveToHead(node *Node) {
	lrucaching.remove(node)
	lrucaching.add(node)
}

// popTail removes the least recently used node (right before tail)
func (lrucaching *LRUCache) popTail() *Node {
	node := lrucaching.tail.prev
	lrucaching.remove(node)
	return node
}

// Get retrieves a value and marks the key as most recently used
func (lrucaching *LRUCache) Get(key int) int {
	if node, found := lrucaching.cache[key]; found {
		lrucaching.moveToHead(node)
		return node.value
	}
	return -1
}

// Put adds a new key or updates an existing one
func (lrucaching *LRUCache) Put(key int, value int) {
	if node, found := lrucaching.cache[key]; found {
		node.value = value
		lrucaching.moveToHead(node)
	} else {
		node := &Node{key: key, value: value}
		lrucaching.cache[key] = node
		lrucaching.add(node)

		if len(lrucaching.cache) > lrucaching.capacity {
			tail := lrucaching.popTail()
			delete(lrucaching.cache, tail.key)
		}
	}
}

func main() {
	// Example test for the LRU Cache
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1

	cache.Put(3, 3)           // Evicts key 2
	fmt.Println(cache.Get(2)) // Output: -1

	cache.Put(4, 4)           // Evicts key 1
	fmt.Println(cache.Get(1)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
