package main

import (
	"fmt"
	"log"
	"os"
)

const metaFileName = "./metadata.txt"

func main() {
	var (
		err                   error
		tmpFile, openMetaFile *os.File
	)

	fmt.Println("Hello")
	if tmpFile, err = os.Create(metaFileName); err != nil {
		log.Fatal(err)
	}

	if err = tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	if openMetaFile, err = os.Open(metaFileName); err != nil {
		log.Fatal(err)
	}

	if err = openMetaFile.Close(); err != nil {
		log.Fatal(err)
	}

	if err = os.Remove(metaFileName); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success")
}
