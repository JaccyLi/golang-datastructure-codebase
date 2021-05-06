package main

import (
	"fmt"

	"suosuoli-golangds/array"
	"suosuoli-golangds/queue"
	"suosuoli-golangds/sort"
	"suosuoli-golangds/stack"
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

func main_stack() {
	ns := stack.NewStack()
	ns.Push(12)
	ns.Push(13)
	ns.Pop()
	d, err := ns.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("data: ", d)
}

func main_queue() {
	nq := queue.New()
	for i := 0; i < 9; i++ {
		nq.Enqueue(i)
	}
	fmt.Println("full: ", nq.IsFull())
	fmt.Printf("%#v\n", nq)
	fmt.Println("peek: ", nq.Peek())
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

func main_tmp() {
	//arr := utils.GenSlice(10, 200)
	//fmt.Println(sort.InsertionSort(arr))
}
