package complexA

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

func RemoveElement() {
	age := map[string]int{
		"alice": 20,
		"tom":   30,
	}
	fmt.Println(age["alice"])
	delete(age, "alice")
	fmt.Println(age["alice"])

}

func RangeMap() {
	ages := map[string]int{
		"alice": 20,
		"tom":   30,
	}
	for name, age := range ages {
		fmt.Println(name, age)
	}
}

func SortMap() {
	var names []string
	ages := map[string]int{
		"alice": 20,
		"tom":   30,
		"bob":   25,
		"peter": 17,
	}
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println(names)
	// sort slice names
	sort.Strings(names)
	for _, v := range names {
		fmt.Printf("%s\t%d\n", v, ages[v])
	}
}

func MapOK() {
	ages := map[string]int{
		"alice": 20,
		"tom":   30,
		"bob":   25,
		"peter": 17,
	}
	_, ok := ages["tim"]
	if !ok {
		fmt.Println(ok)
	}
}

func EqualMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		//
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func MapAsSet() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "debup: %v\n", err)
	}
}

var m = make(map[string]int)

func k(list []string) string {
	fmt.Printf("%T\n", fmt.Sprintf("%q", list))
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
	fmt.Println(m)
}
func Count(list []string) int {

	return m[k(list)]
}
