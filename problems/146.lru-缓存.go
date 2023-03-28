package problems

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start

type LRUCache struct {
	capacity   int
	cacheMap   map[int]*cacheNode
	head, tail *cacheNode
}

type cacheNode struct {
	key, val   int
	prev, next *cacheNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cacheMap: map[int]*cacheNode{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cacheMap[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cacheMap[key]
	if ok {
		node.val = value
		this.moveToHead(node)
	} else {
		node = initNode(key, value)
		this.addToHead(node)
	}
	this.cacheMap[key] = node
	for len(this.cacheMap) > this.capacity {
		this.removeTailNode()
	}
}

func initNode(key, val int) *cacheNode {
	return &cacheNode{
		key: key,
		val: val,
	}
}

func (this *LRUCache) addToHead(node *cacheNode) {
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *cacheNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTailNode() {
	tailNode := this.tail.prev
	tailCacheKey := tailNode.key
	delete(this.cacheMap, tailCacheKey)
	this.removeNode(tailNode)
}

func (this *LRUCache) removeNode(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
