package main

import (
	"fmt"

	"suosuoli-golangds/array"
	"suosuoli-golangds/tree"
)

func main_array() {
	a := array.NewArrayList()
	a.AppendElement(0)
	for i := 1; i < 10; i++ {
		err := a.InsertElement(a.ArraySize, i)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(a.DataStore)
	if err := a.DeleteElement(3); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.DataStore)
	if err := a.DeleteElement(3); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.DataStore)
}

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