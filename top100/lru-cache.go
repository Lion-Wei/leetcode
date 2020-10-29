package top100

type node struct {
	key, val  int
	pre, next *node
}

type LRUCache struct {
	index    map[int]*node
	head     *node
	tail     *node
	capacity int
}

func Constructor(capacity int) LRUCache {
	n := &node{}
	return LRUCache{
		index:    make(map[int]*node, capacity),
		capacity: capacity,
		head:     n,
		tail:     n,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, found := this.index[key]; found {
		if this.head != n {
			n.pre.next = n.next
			n.next.pre = n.pre
			n.next = this.head
			this.head.pre = n
			this.head = n
		}
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, found := this.index[key]; found {
		n.val = value
		n.key = key
		this.Get(key)
		return
	}
	if len(this.index) == this.capacity {
		delete(this.index, this.tail.pre.key)
		this.tail = this.tail.pre
		this.tail.next = nil
	}
	n := &node{
		key:  key,
		val:  value,
		next: this.head,
	}
	this.head.pre = n
	this.head = n
	this.index[key] = n
}
