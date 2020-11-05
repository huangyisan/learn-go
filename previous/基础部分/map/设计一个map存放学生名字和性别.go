package main

import "fmt"

// map[string]map[string]string

func main() {
	studentsMap := make(map[string]map[string]string)
	// 嵌套map都要进行make
	studentsMap["student01"]  = make(map[string]string)
	studentsMap["student01"]["name"] = "tom"
	studentsMap["student01"]["sex"] = "boy"
	studentsMap["student02"]  = make(map[string]string)
	studentsMap["student02"]["name"] = "mary"
	studentsMap["student02"]["sex"] = "girl"

	fmt.Printf("%v",studentsMap)
}
