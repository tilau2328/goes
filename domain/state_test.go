package domain

import "testing"

type test struct {}

func TestNewState(t *testing.T) {
	state := NewState()
	if state == nil {
		t.Errorf("failed to create new state")
	}
}

func TestState_Read(t *testing.T) {
	key := "test"
	value := test{}
	state := NewState()
	state.values[key] = value
	res := state.Read(key)
	if res != value {
		t.Errorf("expected value to be %T but was %T", value, res)
	}
}

func TestState_Write(t *testing.T) {
	key := "test"
	value := test{}
	state := NewState()
	state.Write(key, value)
	res := state.values[key]
	if res != value {
		t.Errorf("expected value to be %T but was %T", value, res)
	}
}