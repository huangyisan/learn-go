package main

// 具体接口

type offCommand struct {
	device device
}

func (o *offCommand) execute() {
	o.device.off()
}
