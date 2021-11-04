package main

import (
    "fmt"
    "suosuoli-golangds/stack"
)

func main() {
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
