package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// var alex person
	// alex.firstname = "Alex"
	// alex.lastname = "Anderson"
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex.anderson@yahoo.com",
			zipCode: 48843,
		},
	}
	// fmt.Println(alex)

	alex.updateName("Jimmy")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstname string) {
	(*pointerToPerson).firstname = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
