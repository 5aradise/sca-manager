package storage

import (
	"context"
	"database/sql"

	"github.com/5aradise/sca-manager/internal/models"
)

const createTargetQuery = `
INSERT INTO targets (mission_id, name, country, notes, is_completed) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, country, notes, is_completed
`

func (q *storage) CreateTarget(ctx context.Context, missionID int32, t models.Target) (models.Target, error) {
	notes := newNullString(t.Notes)
	row := q.db.QueryRowContext(ctx, createTargetQuery,
		missionID,
		t.Name,
		t.Country,
		notes,
		t.IsCompleted,
	)

	err := row.Scan(
		&t.ID,
		&t.Name,
		&t.Country,
		&notes,
		&t.IsCompleted,
	)
	t.Notes = notes.String

	return t, err
}

const deleteTargetByIdQuery = `
DELETE FROM targets WHERE id = $1
`

func (q *storage) DeleteTargetById(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTargetByIdQuery, id)
	return err
}

const updateTargetCompletionByIdQuery = `
UPDATE targets
SET is_completed = $2 
WHERE id = $1
RETURNING id, name, country, notes, is_completed
`

func (q *storage) UpdateTargetCompletionById(ctx context.Context, id int32, isCompleted bool) (models.Target, error) {
	row := q.db.QueryRowContext(ctx, updateTargetCompletionByIdQuery, id, isCompleted)
	var t models.Target
	var notes sql.NullString
	err := row.Scan(
		&t.ID,
		&t.Name,
		&t.Country,
		&notes,
		&t.IsCompleted,
	)
	t.Notes = notes.String
	return t, err
}

const updateTargetNotesByIdQuery = `
UPDATE targets
SET notes = $2 
WHERE id = $1
RETURNING id, name, country, notes, is_completed
`

func (q *storage) UpdateTargetNotesById(ctx context.Context, id int32, notes string) (models.Target, error) {
	nullNotes := newNullString(notes)
	row := q.db.QueryRowContext(ctx, updateTargetNotesByIdQuery, id, nullNotes)
	var t models.Target
	err := row.Scan(
		&t.ID,
		&t.Name,
		&t.Country,
		&nullNotes,
		&t.IsCompleted,
	)
	t.Notes = nullNotes.String
	return t, err
}
