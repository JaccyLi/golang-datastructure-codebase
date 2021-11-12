package oop

import "fmt"

type Human struct {
	Name     string
	Age      int
	Location string
}

type Student struct {
	Human
	Borrow float64
}

type Worker struct {
	Human
	Spent float64
}

func (h *Human) SayHi() {
	fmt.Printf("I'm human, my name is %s, I'm %d years old, I'm live in %s\n",
		h.Name, h.Age, h.Location)
}

func (h *Human) Sing(lyric string) {
	fmt.Printf("%s: %s\n", h.Name, lyric)
}
