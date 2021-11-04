package main

import (
	"fmt"
	"suosuoli-golangds/queue"
)

// =================datastructure main==================== //
func main() {
	nq := queue.New()
	for i := 0; i < 9; i++ {
		nq.Enqueue(i)
	}
	fmt.Println("full: ", nq.IsFull())
	fmt.Printf("%#v\n", nq)
	fmt.Println("peek: ", nq.Peek())
}
