package _1_array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err)
		}
	}
	arr.Print()
	fmt.Println(arr.Len())

	_ = arr.Insert(uint(3), 999)
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err)
		}
	}
	arr.Print()

	_, err := arr.Delete(uint(3))
	if nil != err {
		t.Fatal(err)
	}
	arr.Print()
}

func TestFind(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err)
		}
	}
	arr.Print()

	v, err := arr.Find(uint(3))
	if nil != err {
		t.Fatal(err)
	}
	fmt.Println(v)
}
