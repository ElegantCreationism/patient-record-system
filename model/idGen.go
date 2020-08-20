package model

import (
	"errors"
	"fmt"
	"github.com/sony/sonyflake"
	"log"
)

func GenerateID(personType string) (string, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Printf("flake.NextID() failed with %s\n", err)
		return "", err
	}
	switch personType {
	case "Doctor":
		return fmt.Sprintf("%s-%s", "D", id), nil
	case "Nurse":
		return fmt.Sprintf("%s-%s", "N", id), nil
	case "Patient":
		return fmt.Sprintf("%s-%s", "P", id), nil
	}
	return "", errors.New("failed to create ID: coulld not determine personType")
}
