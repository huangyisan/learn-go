package main

func main() {
	hp := &hp{}
	mac := &mac{}
	mac.setPrinter(hp)
	mac.print()

	epson := &epson{}
	windows := &windows{}
	windows.setPrinter(epson)
	windows.print()
}
