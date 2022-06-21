package main

type mac struct {
	//和printer结合
	printer printer
}

//实现computer接口

func (m *mac) print() {
	m.printer.printFile()
}

func (m *mac) setPrinter(printer printer) {
	m.printer = printer
}
