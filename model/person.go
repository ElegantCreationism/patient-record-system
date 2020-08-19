package model

import "fmt"


type Person struct {
	Name
	Age int
	Gender string
	Address string
	DOB string
}

type Name struct {
	first string
	last string
}

func (n *Name) FullName() {
	fmt.Printf("%s %s", n.first, n.last)
}

func (n *Name) FirstName() {
	fmt.Printf("%s", n.first)
}

func (n *Name) LastName() {
	fmt.Printf("%s", n.last)
}

func NewPerson(age int, firstName, lastName, gender, address, dateOfBirth string) *Person {
	name := newName(firstName, lastName)
	return &Person{
		Name:        *name,
		Age:         age,
		Gender:      gender,
		Address:     address,
		DOB: 		 dateOfBirth,
	}
}

func newName(firstname, lastname string) *Name {
	return &Name{
		first: firstname,
		last:  lastname,
	}
}