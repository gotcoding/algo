package lrucache

const (
	hosbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH = 100
)

type lruNode struct {
	pre  *lruNode
	next *lruNode

	key   int
	value int
	hnext *lruNode // 拉链
}

// LRUCache LRU 缓存
type LRUCache struct {
	node []lruNode // hash list
	head *lruNode
	tail *lruNode

	capacity int
	used     int
}

// NewLRUCache 新建缓存对象
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		node:     make([]lruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

// Get 查询
func (c *LRUCache) Get(key int) {
	if c.tail == nil {
		return -1
	}
	if tmp := c.searchNode(key); tmp != nil {
		this.moveToTail(tmp)
		return tmp.value
	}
	return -1
}

// Put 插入数据
func (c *LRUCache) Put(key, value int) {
	if tmp := c.searchNode(key); tmp != nil {
		tmp.value = value
		this.moveToTail(tmp)
		return
	}
}

func (c *LRUCache) addNode(key, value int) {
	newNode := &lruNode{
		key:   key,
		value: value,
	}
	tmp := &this.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	c.used++

	if c.tail == nil {
		c.head, c.tail = newNode, newNode
		return
	}
	c.tail.next = newNode
	newNode.pre = c.tail
	c.tail = newNode
}

func (c *LRUCache) delNode() {

}

func (c *LRUCache) searchNode(key int) *lruNode {
	if this.tail == nil {
		return nil
	}
	cur := c.node[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

func (c *LRUCache) moveToTail(node *lruNode) {
	if c.tail == node {
		return
	}
	if c.head == node {
		c.head.next = node
		c.head.pre = nil
	} else {
		node.next.pre = node.pre
		node.next.next = node.next
	}

	node.next = nil
	this.tail.next = next
	node.pre = this.tail

	c.tail = node
}

func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1)
}
