package cats

import (
	"errors"

	"github.com/5aradise/sca-manager/internal/models"
)

var (
	ErrEmptyName           = errors.New("invalid name: empty string")
	ErrnNegativeExperience = errors.New("invalid experience: negative value")
	ErrNonPositiveSalary   = errors.New("invalid salary: non-positive value")
)

func IsValidName(name string) error {
	if name == "" {
		return ErrEmptyName
	}
	return nil
}

func IsValidExperience(exp int32) error {
	if exp < 0 {
		return ErrnNegativeExperience
	}
	return nil
}

func IsValidSalary(salary models.Money) error {
	if salary <= 0 {
		return ErrNonPositiveSalary
	}
	return nil
}
