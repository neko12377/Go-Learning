package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
	children struct{
		firstName string
		lastName string
	}
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Parker",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 222,
		},
		children: struct {
			firstName string
			lastName  string
		}{firstName: "chocolate", lastName: "Parker"},
	}

	jim.lastName = "83ik"
	jim.changeName("Jimmy")
	jim.print()
}

func (p *person) changeName (newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print () {
	fmt.Printf("%+v", p)
}
