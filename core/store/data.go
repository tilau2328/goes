package store

type IDataStore interface {
	Delete()
	Update()
	Create()
	List()
	Get()
}

type DataStore struct {
}

func NewDataStore() *DataStore {
	return &DataStore{}
}

func (*DataStore) Get() {

}

func (*DataStore) List() {

}

func (*DataStore) Create() {

}

func (*DataStore) Update() {

}

func (*DataStore) Delete() {

}
