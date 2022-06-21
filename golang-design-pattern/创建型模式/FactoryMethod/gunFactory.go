package main

import "fmt"

// gun工厂函数, 可以通过这个函数, 建造不同的gun

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
