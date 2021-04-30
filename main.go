package main

import (
	"fmt"

	"suosuoli-golangds/array"
	"suosuoli-golangds/sort"
	"suosuoli-golangds/tree"
	"suosuoli-golangds/utils"
)

// =================datastructure main==================== //
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

func main_tree() {
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

// =================datastructure main==================== //

// =================algorithm main==================== //
func main_selectSort() {

	arr := utils.GenSlice(10, 1000)
	fmt.Printf("un: %#v\nsorted:%#v\n", arr, sort.SelectSortAssend(arr))
	fmt.Printf("un: %#v\nsorted:%#v\n", arr, sort.SelectSortDescend(arr))

}

func main() {

}
