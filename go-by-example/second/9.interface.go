package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {

}

type Cat struct {

}

func (dog *Dog) Speak() {
	fmt.Println("wang")
}

func (cat *Cat) Speak() {
	fmt.Println("miao")
}

// 如果要实现迭代对象，使其狗叫猫叫

/*
class Dog(object):
  def __init__(self):
    pass
  def speak(self):
    print("wang")

class Cat(object):
  def __init__(self):
    pass
  def speak(self):
    print("miao")

animal = [Dog(), Cat()]

for i in animal:
  i.speak()
 */

func main() {
	dog := &Dog{}
	cat := &Cat{}

	animal := []Animal{dog, cat}
	for _, v := range animal {
		v.Speak()
	}
}

