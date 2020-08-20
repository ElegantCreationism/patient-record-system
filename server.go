package main

import (
	"fmt"
	"github.com/ElegantCreationism/patient-record-system/model"
)

func main() {
	person := model.NewPerson(25, "bob", "dylan", "Male", "55 Chancery Lane", "!8/06/1997")

	person.FullName()
	fmt.Println()
	person.FirstName()
	fmt.Println()
	person.LastName()
	fmt.Println()
	fmt.Println(person.Age)
	fmt.Println(person.Address)
	fmt.Println(person.Gender)

}
