package model

import "fmt"


type Person struct {
	name
	Age int
	Gender string
	Address string
	DOB string
}

type name struct {
	first string
	last string
}

func (n *name) FullName() {
	fmt.Printf("%s %s", n.first, n.last)
}

func (n *name) FirstName() {
	fmt.Printf("%s", n.first)
}

func (n *name) LastName() {
	fmt.Printf("%s", n.last)
}

func NewPerson(age int, firstName, lastName, gender, address, dateOfBirth string) *Person {
	name := newName(firstName, lastName)
	return &Person{
		name:    *name,
		Age:     age,
		Gender:  gender,
		Address: address,
		DOB:     dateOfBirth,
	}
}

func newName(firstname, lastname string) *name {
	return &name{
		first: firstname,
		last:  lastname,
	}
}