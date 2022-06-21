package main

// iGun 定义一支枪所具备的所有方法.
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
