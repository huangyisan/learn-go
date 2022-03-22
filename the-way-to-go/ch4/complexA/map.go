package complexA

import "fmt"

func GenMap() {
	ages := make(map[string]int)
	age := map[string]int{
		"alice": 20,
		"tom":   30,
	}
	ages["alice"] = 30
	ages["tom"] = 30
	fmt.Println(ages["alice"])
	fmt.Println(age["tom"])
}
