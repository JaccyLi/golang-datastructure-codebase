package array

import (
	"errors"
	"sync"
)

type Array interface {
	Clear()
	String()
	GetSize() int64
	InsertElement(index int64, data interface{}) error
	DeleteElement(index int64) error
	SetElement(index int64) error
	AppendElement(data interface{})
}

type ArrayList struct {
	mu        *sync.Mutex
	ArraySize int64
	DataStore []interface{}
}

func NewArrayList() *ArrayList {
	new := new(ArrayList)
	new.DataStore = make([]interface{}, 0, 20)
	new.ArraySize = 0
	return new
}

func (a *ArrayList) GetSize() int64 {
	return a.ArraySize
}

func (a *ArrayList) AppendElement(data interface{}) {
	a.isFull()
	a.DataStore = a.DataStore[:a.ArraySize+1]
	a.DataStore = append(a.DataStore, data)
	a.ArraySize++
}

func (a *ArrayList) DeleteElement(index int64) error {
	if index < 0 || index > a.ArraySize {
		return errors.New("index out of range")
	}
	a.DataStore = append(a.DataStore[:index], a.DataStore[index+1:])
	a.ArraySize--
	return nil
}

func (a *ArrayList) InsertElement(index int64, data interface{}) error {
	if index < 0 || index > a.ArraySize {
		return errors.New("index out of range")
	}
	// if array is full, double the ram
	a.isFull()
	a.DataStore = a.DataStore[:a.ArraySize+1]
	for i := a.ArraySize; i > index; i-- {
		a.DataStore[i] = a.DataStore[i-1]
	}
	a.DataStore[index] = data
	a.ArraySize++
	return nil
}

func (a *ArrayList) SetElement(index int64, data interface{}) error {
	if index < 0 || index >= a.ArraySize {
		return errors.New("index out of range")
	}
	a.DataStore[index] = data
	return nil
}

func (a *ArrayList) String() {

}

func (a *ArrayList) Clear() {
	a.DataStore = make([]interface{}, 0, 20)
	a.ArraySize = 0
}

func (a *ArrayList) isFull() {
	if a.ArraySize == int64(cap(a.DataStore)) {
		new := make([]interface{}, 2*a.ArraySize)
		copy(new, a.DataStore)
		a.DataStore = new
	}
}
