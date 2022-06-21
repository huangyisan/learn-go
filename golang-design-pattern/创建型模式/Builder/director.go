package main

// 组织房子的建造

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

// 用来切换director
func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
