package service

import (
	"errors"
	"fmt"
	pb "learn-protobuf/pb"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	// save laptop in the store
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

// store laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// return a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// save the laptop in memory
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	// 存入之前判断是否存在了同id的电脑
	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}
	// 使用深拷贝
	other := &pb.Laptop{}
	// 将laptop对象复制到other
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	store.data[other.Id] = other
	return nil
}

// find a laptop by id
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cant copy laptop data: %w", err)
	}
	return other, nil
}
