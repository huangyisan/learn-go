package main

import (

	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)


type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

type FileOpen string

func (f FileOpen) Open(name string) (File,error) {
	file, _ := os.Open(name)
	return file, nil
}


func main() {
	router := gin.Default()

	router.GET("/local/file", func(c *gin.Context) {
		c.File("./tmpdir/LICENSE")
	})

	//
	var fs http.FileSystem = FileOpen("ss")
	//var fs http.FileSystem = http.Dir("./tmpdir")// ...
	//fs = fo.Open("./dd")
	//var fs http.FileSystem = http.Dir("./tmpdir")// ...

	router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("LICENSE", fs)
	})
	router.Run(":7777")
}