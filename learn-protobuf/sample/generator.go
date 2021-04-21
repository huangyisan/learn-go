package sample

import (
	"learn-protobuf/pb"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUbrand()
	name := randomCPUname()
	cpu := &pb.CPU{
		Brand: brand,
		Name:  name,
	}
	return cpu
}
