package design

/*
Company: Amazon(91), Microsoft(19), Facebook(18), Apple(18), Bloomberg(12), VMware(12)

Design and implement a data structure for Least Recently Used (LRU) cache.
It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity,
it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?
*/

// Node ...
type Node struct {
	prev *Node
	next *Node
	val  int
	key  int
}

// LRUCache ...
type LRUCache struct {
	head          *Node
	tail          *Node
	cache         map[int]*Node
	cacheCapacity int
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:         make(map[int]*Node),
		cacheCapacity: capacity,
	}
}

// always add the node as head node
func (this *LRUCache) add(node *Node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) remove(node *Node) {
	if node == nil {
		return
	}
	prev, next := node.prev, node.next
	if prev != nil {
		prev.next = next
	} else {
		this.head = node.next
	}

	if next != nil {
		next.prev = prev
	} else {
		this.tail = node.prev
	}
}

// Get after get the node, remember to move the node to the head
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.add(node)
		return node.val
	}
	return -1
}

// Put ...
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.remove(node)
		this.add(node)
		return
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.add(node)

	// if it is beyond its capacity, then we remove the tail node
	if len(this.cache) > this.cacheCapacity {
		delete(this.cache, this.tail.key)
		this.remove(this.tail)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
