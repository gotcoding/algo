package _1_array

import (
	"errors"
	"fmt"
)

/*
 * 本部分代码只是为了练习数组搬迁
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 */

type Array struct {
	data   []int
	length uint
}

// 新建数组，初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

// 打印数组
func (this *Array) Print() {
	fmt.Println(this.data)
}

// 判断索引是否越界
func (this *Array) IsIndexOutRange(index uint) bool {
	return index >= uint(cap(this.data))
}

// 通过索引查找数组
func (this *Array) Find(index uint) (int, error) {
	if this.IsIndexOutRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

// 指定位置插入
func (this *Array) Insert(index uint, v int) error {
	// 判断是否满或者越界
	if this.Len() == uint(len(this.data)) {
		return errors.New("array is full")
	}
	if index != this.length && this.IsIndexOutRange(index) {
		return errors.New("out of index range")
	}

	for i := this.Len(); i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

// 指定位置删除
func (this *Array) Delete(index uint) (int, error) {
	if this.IsIndexOutRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}
