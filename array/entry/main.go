package main

import (
	"fmt"
	"suosuoli-golangds/array"
)

func main() {
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
