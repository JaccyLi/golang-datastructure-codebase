package array

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

// [1, 2, 3, 4]
func (a *ArrayList) InsertElement(index int64, data interface{}) {
	if a.ArraySize >= int64(cap(a.DataStore)) {
		new := new(ArrayList)
		new.DataStore = make([]interface{}, 0, a.ArraySize*2)
		for i := a.ArraySize; i >= index; i-- {
			a.DataStore[i+1] = a.DataStore[i]
			if i == index {
				a.DataStore[index] = data
			}
		}
		//tmp := a.DataStore[index+1]
		//new.DataStore = append(append(a.DataStore[:index], data), (a.DataStore[index+1:])...)
		//a.DataStore = new.DataStore
		//a.DataStore[index+1] = tmp
		a.ArraySize++
		return
	}
	a.DataStore = append(append(a.DataStore[:index], data), (a.DataStore[index+1:])...)
	a.ArraySize++
}

func (a *ArrayList) AppendElement(data interface{}) {
	a.DataStore = append(a.DataStore, data)
	a.ArraySize++
}

func (a *ArrayList) String() {

}

func (a *ArrayList) Clear() {

}
