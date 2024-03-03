package main

import (
	"container/list"
	"fmt"
)

type myInt int

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	fmt.Println(l.Back().Value)
}
