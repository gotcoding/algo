package arraydev

import (
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	testCases := []struct {
		desc string
		args [][]int
		want []int
	}{
		{
			desc: "多数组合并-空",
			args: [][]int{},
			want: []int{},
		},
		{
			desc: "多数组合并-1个",
			args: [][]int{
				{1, 4, 6, 8},
				{},
			},
			want: []int{1, 4, 6, 8},
		}, {
			desc: "多数组合并-2个",
			args: [][]int{
				{1, 4, 6, 8},
				{3, 5, 9},
			},
			want: []int{1, 3, 4, 5, 6, 8, 9},
		}, {
			desc: "多数组合并-多个",
			args: [][]int{
				{1, 4, 6, 8},
				{3, 5, 9},
				{2, 5, 10},
			},
			want: []int{1, 2, 3, 4, 5, 5, 6, 8, 9, 10},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := MergeSortedArrays(tC.args)
			if len(res) != len(tC.want) {
				t.Error("length not equal")
			}
			for i := 0; i < len(res); i++ {
				if res[i] != tC.want[i] {
					t.Errorf("want: %v , get: %v", tC.want, res)
				}
			}
		})
	}
}
