package model

type GP struct {
	Name string
	Address string
	Doctors []Doctor
}

func NewGP(name, address string) *GP{
	return &GP{
		Name:    name,
		Address: address,
		Doctors: nil,
	}
}