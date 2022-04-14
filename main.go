package main

import (
	"fmt"
	"go-algorithms/stack"
)

func main() {
	var s = stack.StackSlice{}
	s.Push(12)
	s.Push("abc")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
