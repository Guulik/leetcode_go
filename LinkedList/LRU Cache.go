package LinkedList

type CacheDoubleListNode struct {
	Key  int
	Val  int
	Prev *CacheDoubleListNode
	Next *CacheDoubleListNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*CacheDoubleListNode
	//least recently used
	left *CacheDoubleListNode
	//most recently used
	right *CacheDoubleListNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*CacheDoubleListNode, capacity+2),
		left:     &CacheDoubleListNode{Key: 0, Val: 0},
		right:    &CacheDoubleListNode{Key: 0, Val: 0},
	}
	cache.left.Next, cache.right.Prev = cache.right, cache.left
	return cache
}

func (this *LRUCache) listInsert(node *CacheDoubleListNode) {
	rightPrev, right := this.right.Prev, this.right
	right.Prev = node
	rightPrev.Next = node
	node.Prev, node.Next = rightPrev, right
}
func (this *LRUCache) listRemove(node *CacheDoubleListNode) {
	nxt, prv := node.Next, node.Prev
	nxt.Prev, prv.Next = prv, nxt
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		this.listRemove(node)
		this.listInsert(node)
		return this.cache[key].Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		this.listRemove(this.cache[key])
	}
	this.cache[key] = &CacheDoubleListNode{Key: key, Val: value}
	this.listInsert(this.cache[key])

	if len(this.cache) > this.capacity {
		// remove from the list and delete the LRU from hashmap
		lru := this.left.Next
		this.listRemove(lru)
		delete(this.cache, lru.Key)
	}
}
