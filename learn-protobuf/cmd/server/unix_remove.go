package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello")
	metaFileName := "./metadata.txt"
	_, err2 := os.Create(metaFileName)
	if err2 != nil {
		log.Fatal(err2)
	}

	openMetaFile, err := os.Open(metaFileName)
	if err != nil {
		log.Fatal(err)
	}

	err = openMetaFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(metaFileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success")
}
