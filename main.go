package main

import (
	"fmt"

	"suosuoli-golangds/array"
)

func main() {
	a := array.NewArrayList()
	a.AppendElement(20)
	a.AppendElement(250)
	a.AppendElement(340)
	a.AppendElement(309)
	a.AppendElement(600)
	fmt.Println(a.DataStore)
	a.InsertElement(3, 22)
	s := a.GetSize()
	fmt.Println(s)
	fmt.Println(a.DataStore)

	var ia = []int{1, 2, 3, 4, 5, 6, 7, 8}
	x := 0
	data := 999
	b := ia[:x]
	c := ia[x]
	d := ia[x+1:] // x+1
	tmp := ia[x+1]
	fmt.Println("db: ", d)
	new := append(b, c, data)
	fmt.Println("da: ", d)
	fmt.Println("new: ", new)
	new = append(new, d...)
	new[x+2] = tmp

	fmt.Printf("%#v, %#v, %#v, %#v\n", b, c, d, new)
}
