package util

import (
	"math/rand"
	"time"
)

func RandomName() string {
	//var letters = []string{"abc","cc"}
	//fmt.Println(letters)

	var letters = []byte("abccsasdfwerqxcsaer")
	result := make([]byte, 8)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
		//}

	}
	return string(result)
}