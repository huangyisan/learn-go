package main

// 原发器, 原发器来处理快照
type originator struct {
	state string
}

func (e *originator) createMemento() *memento {
	return &memento{
		state: e.state,
	}
}

func (e *originator) restoreMemento(m *memento) {
	e.state = m.getSavedState()
}

func (e *originator) setState(state string) {
	e.state = state
}

func (e *originator) getState() string {
	return e.state
}
