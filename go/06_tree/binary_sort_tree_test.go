package tree

import (
	"testing"
)

func TestBSTInsert(t *testing.T) {
	testCases := []struct {
		desc string
		args []int
		want []int
	}{
		{
			desc: "BST Insert",
			args: []int{62, 58, 88, 47, 99, 73, 37, 35, 51, 93},
			want: []int{35, 37, 47, 51, 58, 62, 73, 88, 93, 99},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var bst *BST
			for _, v := range tC.args {
				bst.Insert(v)
			}
			res := bst.getAll()
			for i, v := range res {
				if tC.want[i] != v {
					t.Errorf("want: %v, result: %v", tC.want, res)
				}
			}
		})
	}
}

func TestBSTDelete(t *testing.T) {
	testCases := []struct {
		desc string
		args []int
		dele []int
		want []int
	}{
		{
			desc: "BST Delete",
			args: []int{62, 58, 88, 47, 99, 73, 37, 35, 51, 93},
			dele: []int{47, 99},
			want: []int{35, 37, 51, 58, 62, 73, 88, 93},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var bst *BST
			for _, v := range tC.args {
				bst.Insert(v)
			}
			for _, v := range tC.dele {
				bst.Delete(v)
			}
			res := bst.getAll()
			for i, v := range res {
				if tC.want[i] != v {
					t.Errorf("want: %v, result: %v", tC.want, res)
				}
			}
		})
	}
}
