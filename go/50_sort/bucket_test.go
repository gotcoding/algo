package sort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	d := []int{4, 2, 9, 1, 8, 3, 5}
	fmt.Println(bucketSort(d))
}
