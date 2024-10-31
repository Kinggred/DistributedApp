package common

import (
	"sync"
	cfg "tic-tac-toe/core/config"
)

type FieldStatus struct {
	arr []int8
	mu  sync.RWMutex
}

func NewFieldStatus() *FieldStatus {
	return &FieldStatus{
		arr: make([]int8, 9),
	}
}

func (fs *FieldStatus) SetByIndex(index int8, value int8) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	valid, value := fs.ValidatePlacement(index, value)
	fs.arr[index] = value

	if valid != true {
		cfg.GameLogger.Printf("Incorrect value provided")
	}
}

func (fs *FieldStatus) GetByIndex(index int8) int8 {
	fs.mu.RLock()
	defer fs.mu.Unlock()
	return fs.arr[index]
}

func (fs *FieldStatus) ValidateValue(value int8) (bool, int8) {
	if value < 3 {
		return true, value
	}
	return false, 0
}

func (fs *FieldStatus) ValidatePlacement(index int8, value int8) (bool, int8) {
	_, value = fs.ValidateValue(value)
	valueOnIndex := fs.GetByIndex(index)
	if valueOnIndex == 0 {
		return true, value
	}
	return false, valueOnIndex
}
