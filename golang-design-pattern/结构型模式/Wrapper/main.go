package main

import "fmt"

func main() {
	pizza := &veggeMania{}

	pizzaWithTomato := &tomatoTopping{
		pizza: pizza,
	}
	pizzaWithCheeseAndTomato := &cheeseTopping{
		pizza: pizzaWithTomato,
	}

	fmt.Println(pizzaWithCheeseAndTomato.getPrice())

}
