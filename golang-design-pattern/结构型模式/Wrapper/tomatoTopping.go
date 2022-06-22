package main

// 加tomato 装饰器
type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	price := t.pizza.getPrice()
	return price + 10
}
