package model

type Hospital struct {
	Name string
	Address string
	Doctors []Doctor
	Nurses []Nurse
}

func NewHospital(name, address string) *Hospital{
	return &Hospital{
		Name:    name,
		Address: address,
		Doctors: nil,
		Nurses:  nil,
	}
}