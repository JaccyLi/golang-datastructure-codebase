package main

import (
	"fmt"
	"suosuoli-golangds/tree"
)

func main() {
	t := tree.Node{Key: 20}
	fmt.Println(t)

	t.Insert(22)
	fmt.Println(t)
	t.Insert(12)
	t.Insert(300)
	t.Insert(5)
	t.Insert(145)
	fmt.Println(t.Right)
	fmt.Println(t.Search(5))
}
