package arraydev

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		arr []int
		len uint
		cap uint
		idx uint
		val int
	}
	testCases := []struct {
		desc string
		args args
		want []int
	}{
		{
			desc: "数组插入1",
			args: args{
				arr: []int{},
				len: 0,
				cap: 5,
				idx: 2,
				val: 90,
			},
			want: []int{0, 0, 90, 0, 0},
		}, {
			desc: "数组插入2",
			args: args{
				arr: []int{},
				len: 0,
				cap: 5,
				idx: 0,
				val: 90,
			},
			want: []int{90, 0, 0, 0, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := NewArray(tC.args.cap)
			copy(a.data, tC.args.arr)
			a.length = tC.args.len
			a.Insert(tC.args.idx, tC.args.val)
			if !IsEqual(a.data, tC.want) {
				t.Errorf("args: %v, want: %v, got: %v", tC.args, tC.want, a.data)
			}
		})
	}
}

func IsEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) || cap(arr1) != cap(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
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
