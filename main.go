package main

import (
	"container/list"
	"fmt"

	"helloGo/itertools"
)

func main() {
	items := list.New()
	for i := 0; i < 5; i++ {
		items.PushBack(i)
	}

	iter := itertools.New(items)
	for iter.Next() {
		item := iter.Item()
		fmt.Println(item.Value)
	}
}
