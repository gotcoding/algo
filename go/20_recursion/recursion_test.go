package recursion

import (
	"fmt"
	"testing"
	"time"
)

func TestTimesRec(t *testing.T) {
	n := 40
	start := time.Now()
	fmt.Println(timesRec(n))
	fmt.Println(time.Since(start))
}

func TestFabi(t *testing.T) {
	n := 40
	f := NewFibs(n)
	start := time.Now()
	fmt.Println(timesRec(n))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(f.fibonacciDic(n))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(f.fibonacciArr(n))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(inFor(n))
	fmt.Println(time.Since(start))
}
