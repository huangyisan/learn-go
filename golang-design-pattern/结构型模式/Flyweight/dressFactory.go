package main

import "fmt"

// 享元工厂

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	// 如果判断dress存在,则直接返回得到的dress
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	// 没有,则往dress map里面追加
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	// 没有,则往dress map里面追加
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
