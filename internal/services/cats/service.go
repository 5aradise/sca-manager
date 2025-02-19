package cats

import (
	"context"

	"github.com/5aradise/sca-manager/internal/models"
)

type (
	catStorage interface {
		CreateCat(ctx context.Context, cat models.Cat) (models.Cat, error)
		GetCatById(ctx context.Context, id int32) (models.Cat, error)
		ListCats(ctx context.Context) ([]models.Cat, error)
		UpdateCatSalaryById(ctx context.Context, id int32, salary models.Money) (models.Cat, error)
		DeleteCatById(ctx context.Context, id int32) error
	}

	breedValidator interface {
		IsValidBreed(breed string) error
	}

	service struct {
		stor catStorage
		v    breedValidator
	}
)

func New(s catStorage, v breedValidator) *service {
	return &service{s, v}
}

func (s *service) CreateCat(ctx context.Context, cat models.Cat) (models.Cat, error) {
	if err := IsValidName(cat.Name); err != nil {
		return models.Cat{}, err
	}

	if err := IsValidExperience(cat.YearsOfExperience); err != nil {
		return models.Cat{}, err
	}

	if err := IsValidSalary(cat.Salary); err != nil {
		return models.Cat{}, err
	}

	if err := s.v.IsValidBreed(cat.Breed); err != nil {
		return models.Cat{}, err
	}

	newCat, err := s.stor.CreateCat(ctx, cat)
	if err != nil {
		return models.Cat{}, err
	}

	return newCat, nil
}

func (s *service) GetCat(ctx context.Context, id int32) (models.Cat, error) {
	return s.stor.GetCatById(ctx, id)
}

func (s *service) ListCats(ctx context.Context) ([]models.Cat, error) {
	return s.stor.ListCats(ctx)
}

func (s *service) UpdateCatSalary(ctx context.Context, id int32, salary models.Money) (models.Cat, error) {
	if err := IsValidSalary(salary); err != nil {
		return models.Cat{}, err
	}

	return s.stor.UpdateCatSalaryById(ctx, id, salary)
}

func (s *service) DeleteCat(ctx context.Context, id int32) error {
	return s.stor.DeleteCatById(ctx, id)
}
