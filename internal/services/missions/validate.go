package missions

import (
	"errors"
	"fmt"

	"github.com/5aradise/sca-manager/internal/models"
)

const (
	MinTargetsCount = 1
	MaxTargetsCount = 3
)

var (
	ErrInvalidTargetsCount = errors.New("invalid targets: wrong number (minimum: 1, maximum: 3)")
	ErrEmptyName           = errors.New("invalid name: empty string")
	ErrEmptyCountry        = errors.New("invalid country: empty string")
)

func IsValidTargets(targets []models.Target) error {
	tsCount := len(targets)
	if !(MinTargetsCount <= tsCount && tsCount <= MaxTargetsCount) {
		return ErrInvalidTargetsCount
	}

	for _, target := range targets {
		if err := IsValidTarget(target); err != nil {
			return fmt.Errorf("invalid targets: %w", err)
		}
	}

	return nil
}

func IsValidTarget(target models.Target) error {
	if target.Name == "" {
		return ErrEmptyName
	}

	if target.Country == "" {
		return ErrEmptyCountry
	}

	return nil
}
