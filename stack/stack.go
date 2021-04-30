package stack

import (
	"errors"
)

type StackBehavior interface {
	Push(interface{}) error
	Pop() interface{}
	IsEmpty() bool
	IsFull() bool
	Peek() interface{}
}

type Stack struct {
	Size      int64
	Top       interface{}
	DataStore []interface{}
}

func NewStack() *Stack {
	s := new(Stack)
	s.DataStore = make([]interface{}, 0, 10)
	s.Size = 0
	return s
}

func (s *Stack) Push(data interface{}) error {
	if IsFull(s) {
		return errors.New("stack overflow")
	}
	s.DataStore = append(s.DataStore, data)
	s.Size++
	return nil
}

func Pop(s *Stack) error {
	if IsEmpty(s) {
		return errors.New("stack is empty")
	}
	s.DataStore = s.DataStore[:len(s.DataStore)-1]
	s.Size--
	return nil
}

func IsEmpty(s *Stack) bool {
	return s.Size == 0
}

func IsFull(s *Stack) bool {
	return s.Size == int64(cap(s.DataStore))
}
