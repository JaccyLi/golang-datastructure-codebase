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
	if s.IsFull() {
		return errors.New("stack overflow")
	}
	s.DataStore = append(s.DataStore, data)
	s.Size++
	return nil
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return errors.New("stack is empty")
	}
	s.DataStore = s.DataStore[:len(s.DataStore)-1]
	s.Size--
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) IsFull() bool {
	return s.Size == int64(cap(s.DataStore))
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	rData := s.DataStore[len(s.DataStore)-1:]
	return rData, nil
}
