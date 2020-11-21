package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	testCases := []struct {
		desc string
		args args
		want bool
	}{
		{
			desc: "二分查找",
			args: args{
				arr: []int{2, 3, 3, 3, 5, 8, 9, 11, 23, 54, 65, 56},
				val: 3,
			},
			want: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			b := binarySearch(tC.args.arr, tC.args.val)
			if b != tC.want {
				t.Errorf("arr: %+v, search: %d, get: %v, want: %v", tC.args.arr,
					tC.args.val, b, tC.want)
			}
		})
	}
}
