package main

import "fmt"

// 使用for-range方式遍历

func main() {
	countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}

	for k,v := range countryCapitalMap {
		fmt.Println(k,v)
	}


//	 使用for-range遍历多维map
	studentsMap := make(map[string]map[string]string)
	// 嵌套map都要进行make
	studentsMap["student01"]  = make(map[string]string)
	studentsMap["student01"]["name"] = "tom"
	studentsMap["student01"]["sex"] = "boy"
	studentsMap["student02"]  = make(map[string]string)
	studentsMap["student02"]["name"] = "mary"
	studentsMap["student02"]["sex"] = "girl"

	for _,v := range studentsMap {
		for k2, v2 :=range v {
			fmt.Println(k2,v2)
		}
	}
}

