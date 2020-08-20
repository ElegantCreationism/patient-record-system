package model

// All Patient Id's start with a 'P'
type Patient struct {
	ID string
	Person
	Address string
	AssignedDoctor string
	RegisteredHospital Hospital
	RegisteredGP GP
}


func NewPatient(person Person, address string, assignedDoctor string, registeredHospital string, RegisteredGP string)  *Patient {
	id, err := GenerateID("Nurse")
	if err != nil{
		return nil
	}



	return &Patient{
		ID:                 id,
		Person:             person,
		Address:            address,
		AssignedDoctor:     assignedDoctor,
		RegisteredHospital: registeredHospital,
		RegisteredGP:       RegisteredGP,
	}
}