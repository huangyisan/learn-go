package service

import (
	pb "learn-protobuf/pb"
	"sync"
)

type LaptopStore interface {
	// save laptop in the store
	Save(laptop *pb.Laptop) error
}

// store laptop in memory
type InMemoryStore struct {
	mutex sync.Mutex
	data  map[string]*pb.Laptop
}

func NewInMemory