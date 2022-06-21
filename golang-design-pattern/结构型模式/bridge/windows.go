package main

type windows struct {
	//和printer结合
	printer printer
}

//实现computer接口

func (w *windows) print() {
	w.printer.printFile()
}

func (w *windows) setPrinter(printer printer) {
	w.printer = printer
}
