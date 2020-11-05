package hashmap

import "fmt"

const (
	expandFactor = 0.75
)

type KeyPairs struct {
	key   string
	value interface{}
	next  *KeyPairs
}

type HashMap struct {
	m   []*KeyPairs // hash表的桶
	cap int
	len int
}

// NewHashMap 新建散列表
func NewHashMap(cap int) *HashMap {
	if cap < 16 {
		cap = 16
	}
	return &HashMap{
		m:   make([]*KeyPairs, cap, cap),
		cap: cap,
		len: 0,
	}
}

// BKDRHash 哈希值
func BKDRHash(str string, cap int) int {
	seed := int(131) // 31 131 1313 13131 131313 etc..
	hash := int(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int(str[i])
	}
	return hash % cap
}

// Add 存
func (hm *HashMap) Add(key string, value interface{}) {
	index := BKDRHash(key, hm.cap)
	cur := hm.m[index]
	hm.len++
	// 如果不存在则插入
	if cur == nil {
		hm.m[index] = &KeyPairs{key, value, nil}
	} else {
		// 如果存在则更新
		for cur.next != nil {
			if cur.key == key {
				cur.value = value
				return
			}
			cur = cur.next
		}
		cur.next = &KeyPairs{key, value, nil}
	}
	// 扩容
	if float64(hm.len)/float64(hm.cap) > expandFactor {
		newH := NewHashMap(2 * hm.cap)
		for _, pair := range hm.m {
			if pair == nil {
				continue
			}
			newH.Add(pair.key, pair.value)
			for pair.next != nil {
				pair = pair.next
				newH.Add(pair.key, pair.value)
			}
		}
		hm.cap = newH.cap
		hm.m = newH.m
	}
}

// Get 查询
func (hm *HashMap) Get(key string) interface{} {
	index := BKDRHash(key, hm.cap)
	cur := hm.m[index]
	if cur == nil {
		return nil
	} else if cur.key == key {
		return cur.value
	}
	for cur.next != nil {
		if cur.key == key {
			return cur.value
		}
		cur = cur.next
	}
	return nil
}

// Print 打印
func (hm *HashMap) Print() {
	for _, pairs := range hm.m {
		if pairs == nil {
			continue
		}
		index := BKDRHash(pairs.key, hm.cap)
		fmt.Printf("index: %+v ", index)
		fmt.Printf("key:%v, value:%v ", pairs.key, pairs.value)
		for pairs.next != nil {
			pairs = pairs.next
			fmt.Printf("key:%v, value:%v ", pairs.key, pairs.value)
		}
		fmt.Printf("\n")
	}
}

// Delete 删除
