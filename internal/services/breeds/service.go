package breeds

import (
	"errors"
	"time"

	"github.com/5aradise/sca-manager/pkg/types"
)

var (
	ErrBreedUnfound = errors.New("invalid breed: not found in register")
)

type service struct {
	breeds types.Set[string]
}

func New(apiKey string, reqTimeout time.Duration) (*service, error) {
	breeds, err := listBreeds(apiKey, reqTimeout)
	if err != nil {
		return nil, err
	}
	return &service{breeds}, nil
}

func (s *service) IsValidBreed(breed string) error {
	if !s.breeds.Has(breed) {
		return ErrBreedUnfound
	}
	return nil
}
