package queue

import "sync"

type QueueBehavior interface {
	Enqueue(interface{})
	Dequeue()
	IsEmpty() bool
	IsFull() bool
	Peek() interface{}
}

type Queue struct {
	DataStore []interface{}
	Len       int
	Lock      *sync.Mutex
}

func New() *Queue {
	newq := &Queue{}
	newq.DataStore = make([]interface{}, 0, 10)
	newq.Len = 0
	newq.Lock = new(sync.Mutex)
	return newq
}

func (q *Queue) Enqueue(data interface{}) {
	q.DataStore = append(q.DataStore, data)
	q.Len++
}

func (q *Queue) Dequeue() (re interface{}) {
	if !q.IsEmpty() {
		re, q.DataStore = q.DataStore[0], q.DataStore[1:]
		q.Len--
	}
	return
}

func (q *Queue) IsEmpty() bool {
	q.Lock.Lock()
	defer q.Lock.Unlock()
	return q.Len == 0
}

func (q *Queue) IsFull() bool {
	return q.Len == cap(q.DataStore)
}

func (q *Queue) Peek() interface{} {
	return q.DataStore[0]
}
