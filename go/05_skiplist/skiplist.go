package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MAX_LEVEL = 16 //最高层数
)

// Node 跳表结构体
type Node struct {
	v        interface{} //跳表保存的值
	score    int         // 用于排序的分值
	level    int         // 层高
	forwards []*Node     //每层前进的指针
}

// NewSkipListNode 新建跳表节点
func NewSkipListNode(v interface{}, score, level int) *Node {
	return &Node{v, score, level, make([]*Node, level, level)}
}

// SkipList 跳表结构体
type SkipList struct {
	head   *Node // 跳表头结点
	level  int   // 跳表当前层数
	length int   // 跳表长度
}

// NewSkipList 新建跳表
func NewSkipList() *SkipList {
	head := NewSkipListNode(0, math.MaxInt32, MAX_LEVEL)
	return &SkipList{head, 1, 0}
}

// Insert 插入
func (sl *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找插入位置
	cur := sl.head
	//记录每层的路径
	update := [MAX_LEVEL]*Node{}
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点
	newNode := NewSkipListNode(v, score, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

// Find 查找
func (sl *SkipList) Find(v interface{}, score int) *Node {
	if nil == v || sl.length == 0 {
		return nil
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

// Print 打印跳表
func (sl *SkipList) Print() {
	fmt.Printf("level: %+v, length: %+v\n", sl.level, sl.length)
}
