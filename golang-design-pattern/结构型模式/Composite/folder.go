package main

import "fmt"

type folder struct {
	compoents []component
	name      string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.compoents {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.compoents = append(f.compoents, c)
}
