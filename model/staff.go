package model

// All doctor Id's start with a 'D'
type Doctor struct {
	ID string
	Person
	Specialisation string
	PlaceOfWork Hospital
	OnDuty bool
}

// All nurse Id's start with an 'N'
type Nurse struct {
	ID string
	Person
	Ward string
	PlaceOfWork Hospital
	OnDuty bool
}

func NewDoctor(person Person, specialisation string, placeOfWork Hospital)  *Doctor {
	id, err := GenerateID("Doctor")
	if err != nil{
		return nil
	}
	return &Doctor{
		ID:             id,
		Person:         person,
		Specialisation: specialisation,
		PlaceOfWork:    placeOfWork,
		OnDuty:         false,
	}
}

func NewNurse(person Person, ward string, placeOfWork Hospital)  *Nurse {
	id, err := GenerateID("Nurse")
	if err != nil{
		return nil
	}
	return &Nurse{
		ID:             id,
		Person:         person,
		Ward: 			ward,
		PlaceOfWork:    placeOfWork,
		OnDuty:         false,
	}
}