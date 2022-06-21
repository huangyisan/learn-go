package main

import "fmt"

func printGunDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

func main() {
	ak47, _ := getGun("ak47")
	printGunDetails(ak47)
}
