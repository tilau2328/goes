package store

import "testing"

func TestNewDataStore(t *testing.T) {
	dataStore := NewDataStore()
	if dataStore == nil {
		t.Errorf("failed to create data store")
	}
}

func TestDataStore_Get(t *testing.T) {

}

func TestDataStore_List(t *testing.T) {

}

func TestDataStore_Create(t *testing.T) {

}

func TestDataStore_Update(t *testing.T) {

}

func TestDataStore_Delete(t *testing.T) {

}
