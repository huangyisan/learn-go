package main

// gun的具体产品,其实是抽象了枪支
// gun的属性
type gun struct {
	name  string
	power int
}

// gun 实现iGun接口
func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}
