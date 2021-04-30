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
	// search for laptop with filter, returns one by one via the found Function
	Search(filter *pb.Filter, found func(laptop *pb.Laptop) error) error
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
	// other := &pb.Laptop{}
	// 将laptop对象复制到other
	other, err := deepCopy(laptop)
	// err := copier.Copy(other, laptop)
	if err != nil {
		return err
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
	// other := &pb.Laptop{}
	return deepCopy(laptop)
	// other, err := deepCopy(laptop)
	// if err != nil {
	// 	return nil, fmt.Errorf("cant copy laptop data: %w", err)
	// }
	// return other, nil
}

// search
func (store *InMemoryLaptopStore) Search(
	filter *pb.Filter, found func(laptop *pb.Laptop) error) error {
	// 获取读锁
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, laptop := range store.data {
		if isQualified(filter, laptop) {
			// deep copy
			other, err := deepCopy(laptop)
			if err != nil {
				return err
			}
			err = found(other)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}
	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}
	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}
	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}
	return true
}

func toBit(memory *pb.Memory) uint64 {
	value := memory.GetValue()
	switch memory.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3
	case pb.Memory_KILOBYTE:
		return value << 13
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}

func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cant copy laptop data: %w", err)
	}
	return other, nil
}
