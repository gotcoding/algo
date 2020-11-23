package lc20

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	fmt.Println(isValid(s))
	s1 := "()[){}"
	fmt.Println(isValid(s1))
}
