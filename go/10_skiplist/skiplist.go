package skiplist

import (
	"fmt"
	"math/rand"
)

//Node 跳表节点
type Node struct {
	index     uint64      // 节点索引
	data      interface{} // 节点值
	nextNodes []*Node     // 节点每层的后续指针，塔
}

// newNode 	生成新的节点
func newNode(index uint64, v interface{}, level int) *Node {
	return &Node{index, v, make([]*Node, level, level)}
}

// SkipList 跳表结构体
type SkipList struct {
	level  int   // 层
	length int32 // 长度
	head   *Node // 头节点
	tail   *Node // 尾节点
}

// NewSkipList 新建跳表
func NewSkipList(level int) *SkipList {
	head := newNode(0, nil, level)
	var tail *Node
	for i := 0; i < level; i++ {
		head.nextNodes[i] = tail
	}
	return &SkipList{
		level:  level,
		length: 0,
		head:   head,
		tail:   tail,
	}
}

// Print 打印跳表
func (sl *SkipList) Print() {
	// 自上而下逐层遍历
	for i := sl.level - 1; i >= 0; i-- {
		fmt.Print("head->")
		levelNode := sl.head.nextNodes[i]
		levelZero := sl.head.nextNodes[0]
		// 从head到tail遍历
		for j := 0; j < int(sl.length); j++ {
			if levelNode != nil && levelZero.index == levelNode.index {
				fmt.Printf("%10s->", levelZero.data)
				levelNode = levelNode.nextNodes[i]
			} else {
				fmt.Print("            ")
			}
			levelZero = levelZero.nextNodes[0]
		}
		fmt.Printf("tail\n")
	}
}

// 查找匹配的节点
// 返回第一个参数是前序节点的切片
// 返回第二个参数是后续节点
func (sl *SkipList) searchWithPreviousNodes(index uint64) ([]*Node, *Node) {
	previous := make([]*Node, sl.level)
	cur := sl.head

	for l := sl.level - 1; l >= 0; l-- {
		for cur.nextNodes[l] != sl.tail && cur.nextNodes[l].index < index {
			cur = cur.nextNodes[l]
		}
		previous[l] = cur
	}
	if cur.nextNodes[0] != sl.tail {
		cur = cur.nextNodes[0]
	}
	return previous, cur
}

// Search 查找
func (sl *SkipList) Search(index uint64) *Node {
	cur := sl.head
	for l := sl.level - 1; l >= 0; l-- {
		for cur.nextNodes[l] != sl.tail && cur.nextNodes[l].index < index {
			cur = cur.nextNodes[l]
		}
	}
	cur = cur.nextNodes[0]
	if cur == sl.tail || cur.index > index {
		return nil
	} else if cur.index == index {
		return cur
	} else {
		return nil
	}
}

// Insert 插入
func (sl *SkipList) Insert(index uint64, v interface{}) {
	// 查找插入的点
	pres, curNode := sl.searchWithPreviousNodes(index)
	// 如果索引相等，则更新
	if curNode.index == index && curNode != sl.head {
		curNode.data = v
		return
	}
	rl := sl.randomLevel()
	// 构造节点
	newNode := newNode(index, v, rl)
	// 调整指针
	for l := rl - 1; l >= 0; l-- {
		newNode.nextNodes[l] = pres[l].nextNodes[l]
		pres[l].nextNodes[l] = newNode
		pres[l] = nil
	}
	for l := rl; l < sl.level; l++ {
		pres[l] = nil
	}
	sl.length++
}

// Delete 删除
func (sl *SkipList) Delete(index uint64) {
	// 查找删除的节点
	pres, curNode := sl.searchWithPreviousNodes(index)
	// 节点存在则删除
	if curNode.index == index && curNode != sl.head {
		// 调整指针
		for l := len(curNode.nextNodes) - 1; l >= 0; l-- {
			pres[l].nextNodes[l] = curNode.nextNodes[l]
			pres[l] = nil
		}
	}
	for l := len(curNode.nextNodes); l < sl.level; l++ {
		pres[l] = nil
	}
	sl.length--
}

// 获取随机层数
func (sl *SkipList) randomLevel() int {
	level := 1
	// for rand.Float64() < PROBABILITY && level < s.level {
	// 	level++
	// }
	for i := 1; i < sl.level; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}
	return level
}
