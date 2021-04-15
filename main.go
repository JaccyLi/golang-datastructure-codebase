package main

import (
	"fmt"

	"suosuoli-golangds/array"
)

func main() {
	a := array.NewArrayList()
	a.AppendElement(20)
	a.AppendElement(200)
	//a.InsertElement(0, 22)
	s := a.GetSize()
	fmt.Println(s)
	fmt.Printf("%#v\n", a.DataStore)
}
