package main

// 负责人
type careTaker struct {
	mementoArray []*memento
}

func (c *careTaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *careTaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}
