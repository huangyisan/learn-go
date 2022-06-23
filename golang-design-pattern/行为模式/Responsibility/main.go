package main

func main() {
	doctor := &doctor{}
	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)
	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
