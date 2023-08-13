package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
func (pointer *person) updateName(newfirstname string) {
	(*pointer).first_name = newfirstname
}

func main() {
	// In this type of Intialzation the order of fields matters
	// kinaan := person{"Kinaan Aamir", "Khan"}
	// In this Initialzation we do not use the order
	// kinaan := person{first_name: "Kinaan Aamir", last_name: "Khan"}
	// Definiation of kinaan with person. Go assigns zero values to each field inside the struct
	// for string "", int and float is 0, bool is false
	// var kinaan person
	// kinaan.first_name = "Kinaan Aamir"
	// kinaan.last_name = "Khan"
	// fmt.Println(kinaan)
	// You can use printf to display all field names and values
	// fmt.Printf("%+v", kinaan)

	// var jim person
	//jim = person{
	//	first_name: "Jim",
	//	last_name:  "Anderson",
	//	contact:    contactInfo{email: "jim@gmail.com", zipCode: 65125},
	//}
	//fmt.Printf("%+v", jim)

	var jim person
	jim = person{
		first_name:  "Jim",
		last_name:   "Anderson",
		contactInfo: contactInfo{email: "jim@gmail.com", zipCode: 65125},
	}
	jim.print()
	fmt.Println("\n------")
	//var jimPointer *person
	//jimPointer := &jim

	jim.updateName("kinaan")
	jim.print()
}
