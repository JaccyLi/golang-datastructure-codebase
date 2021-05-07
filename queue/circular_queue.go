package queue

type CircularQueueBehavior interface {
	IsEmpty() bool
	IsFull() bool
	Enqueue(interface{}) bool
	Dequeue() interface{}
	Peek() interface{}
	String() []interface{}
}

type CircularQueue struct {
	DataStore []interface{}
	Capacity  int
	HeadIdx   int
	RearIdx   int
}

func NewCQ(len int) *CircularQueue {
	if len == 0 {
		return nil
	}
	return &CircularQueue{
		DataStore: make([]interface{}, 0, len),
		Capacity:  len,
		HeadIdx:   0,
		RearIdx:   0,
	}
}

func (cq *CircularQueue) IsFull() bool {
	if (cq.RearIdx+1)%cq.Capacity == cq.HeadIdx {
		return true
	}
	return false
}

func (cq *CircularQueue) IsEmpty() bool {
	if cq.RearIdx == cq.HeadIdx {
		return true
	}
	return false
}

func (cq *CircularQueue) Enqueue(data interface{}) bool {
	if cq.IsFull() {
		return false
	}
	cq.DataStore[cq.RearIdx] = data
	cq.RearIdx = (cq.RearIdx + 1) % cq.Capacity
	return true
}

func (cq *CircularQueue) Dequeue() (re interface{}) {
	if cq.IsEmpty() {
		return false
	}
	re = cq.DataStore[cq.HeadIdx]
	cq.HeadIdx = (cq.HeadIdx + 1) % cq.Capacity
	return
}

func (cq *CircularQueue) Peek() interface{} {
	return cq.DataStore[cq.HeadIdx]
}


func (cq *CircularQueue) String() []interface{} {

	return nil
}