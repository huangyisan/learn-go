package main

// 具体接口

type onCommand struct {
	device device
}

func (o *onCommand) execute() {
	o.device.on()
}