package main

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("patient doctor check done")
		d.next.execute(p)
	}
	fmt.Println("finished")

}

func (d *doctor) setNext(next department) {
	d.next = next
}
