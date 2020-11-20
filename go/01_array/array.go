package arraydev

import (
	"errors"
	"fmt"
)

/*
 * 实现一个固定长度的数组
 * 本部分代码只是为了练习数组搬迁
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 */

type Array struct {
	data   []int
	length uint
}

// NewArray 新建数组，初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

func (a *Array) Len() uint {
	return a.length
}

// Print 打印数组
func (a *Array) Print() {
	fmt.Println(a.data)
}

// IsIndexOutRange 判断索引是否越界
func (a *Array) IsIndexOutRange(index uint) bool {
	return index >= uint(cap(a.data))
}

// Find 通过索引查找数组
func (a *Array) Find(index uint) (int, error) {
	if a.IsIndexOutRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

// Insert 指定位置插入
func (a *Array) Insert(index uint, v int) error {
	// 判断是否满或者越界
	if a.Len() == uint(len(a.data)) {
		return errors.New("array is full")
	}
	if index != a.length && a.IsIndexOutRange(index) {
		return errors.New("out of index range")
	}

	for i := a.Len(); i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

// Delete 指定位置删除
func (a *Array) Delete(index uint) (int, error) {
	if a.IsIndexOutRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}
