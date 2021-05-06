package queue

type CircularQueueBehavior interface {
	Enqueu(interface{})
	Dequeu() interface{}
	Peek() interface{}
}

type CircularQueue struct {
	Len       int
	DataStore []interface{}
}

func NewCq(len int) *CircularQueue {
	ncq := &CircularQueue{}
	ncq.DataStore = make([]interface{}, 0, len)
	ncq.Len = 0
	return ncq
}

func (cq *CircularQueue) Enqueu(data interface{}) {
	if !cq.IsFull() {
		cq.DataStore = append(cq.DataStore, data)
		cq.Len++
	}
}

func (cq *CircularQueue) Dequeu() (re interface{}) {
	if !cq.IsEmpty() {
		re, cq.DataStore = cq.DataStore[0], cq.DataStore[1:]
	}
	return
}

func (cq *CircularQueue) IsFull() bool {
	return cq.Len+1 == cap(cq.DataStore)
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.Len == 0
}

func (cq *CircularQueue) Peek() interface{} {
	return cq.DataStore[0]
}
