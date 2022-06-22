package main

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	price := c.pizza.getPrice()
	return price + 20
}
