package storage

import (
	"context"

	"github.com/5aradise/sca-manager/internal/models"
)

const createCatQuery = `
INSERT INTO cats (name, years_of_experience, breed, salary_in_cents) 
VALUES ($1, $2, $3, $4)
RETURNING id, name, years_of_experience, breed, salary_in_cents
`

func (q *storage) CreateCat(ctx context.Context, c models.Cat) (models.Cat, error) {
	row := q.db.QueryRowContext(ctx, createCatQuery,
		c.Name,
		c.YearsOfExperience,
		c.Breed,
		c.Salary.Cents(),
	)
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.YearsOfExperience,
		&c.Breed,
		&c.Salary,
	)
	return c, err
}

const deleteCatByIdQuery = `
DELETE FROM cats WHERE id = $1
`

func (q *storage) DeleteCatById(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCatByIdQuery, id)
	return err
}

const getCatByIdQuery = `
SELECT id, name, years_of_experience, breed, salary_in_cents FROM cats 
WHERE id = $1
`

func (q *storage) GetCatById(ctx context.Context, id int32) (models.Cat, error) {
	row := q.db.QueryRowContext(ctx, getCatByIdQuery, id)
	var c models.Cat
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.YearsOfExperience,
		&c.Breed,
		&c.Salary,
	)
	return c, err
}

const listCatsQuery = `
SELECT id, name, years_of_experience, breed, salary_in_cents FROM cats
`

func (q *storage) ListCats(ctx context.Context) ([]models.Cat, error) {
	rows, err := q.db.QueryContext(ctx, listCatsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cs []models.Cat
	for rows.Next() {
		var c models.Cat
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.YearsOfExperience,
			&c.Breed,
			&c.Salary,
		); err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cs, nil
}

const updateCatSalaryByIdQuery = `
UPDATE cats 
SET salary_in_cents = $2 
WHERE id = $1
RETURNING id, name, years_of_experience, breed, salary_in_cents
`

func (q *storage) UpdateCatSalaryById(ctx context.Context, id int32, salary models.Money) (models.Cat, error) {
	row := q.db.QueryRowContext(ctx, updateCatSalaryByIdQuery, id, salary.Cents())
	var c models.Cat
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.YearsOfExperience,
		&c.Breed,
		&c.Salary,
	)
	return c, err
}
