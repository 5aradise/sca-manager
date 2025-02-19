package storage

import (
	"context"
	"database/sql"

	"github.com/5aradise/sca-manager/internal/models"
)

const createMissionQuery = `
INSERT INTO missions (cat_id, is_completed) 
VALUES ($1, $2)
RETURNING id, cat_id, is_completed
`

func (q *storage) CreateMission(ctx context.Context, m models.Mission) (models.Mission, error) {
	catID := newNullInt32(m.CatID)
	row := q.db.QueryRowContext(ctx, createMissionQuery, catID, m.IsCompleted)
	err := row.Scan(&m.ID, &catID, &m.IsCompleted)
	m.CatID = catID.Int32
	return m, err
}

const deleteMissionById = `
DELETE FROM missions WHERE id = $1
`

func (q *storage) DeleteMissionById(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteMissionById, id)
	return err
}

const getMissionByIdQuery = `
SELECT m.id, m.cat_id, m.is_completed,
    t.id AS target_id, t.name, t.country, t.notes, t.is_completed
FROM missions m
LEFT JOIN targets t ON m.id = t.mission_id
WHERE m.id = $1
`

func (q *storage) GetMissionById(ctx context.Context, id int32) (models.Mission, error) {
	rows, err := q.db.QueryContext(ctx, getMissionByIdQuery, id)
	if err != nil {
		return models.Mission{}, err
	}
	defer rows.Close()
	var m models.Mission
	var catID sql.NullInt32
	var notes sql.NullString

	var rowCount int
	for rows.Next() {
		rowCount++
		var t models.Target
		if err := rows.Scan(
			&m.ID,
			&catID,
			&m.IsCompleted,
			&t.ID,
			&t.Name,
			&t.Country,
			&notes,
			&t.IsCompleted,
		); err != nil {
			return models.Mission{}, err
		}
		m.CatID = catID.Int32
		t.Notes = notes.String
		m.Targets = append(m.Targets, t)
	}
	if rowCount == 0 {
		return models.Mission{}, sql.ErrNoRows
	}

	if err := rows.Close(); err != nil {
		return models.Mission{}, err
	}
	if err := rows.Err(); err != nil {
		return models.Mission{}, err
	}
	return m, nil
}

const listMissionsQuery = `
SELECT m.id, m.cat_id, m.is_completed,
       t.id AS target_id, t.name, t.country, t.notes, t.is_completed
FROM missions m
LEFT JOIN targets t ON m.id = t.mission_id
ORDER BY m.id
`

func (q *storage) ListMissions(ctx context.Context) ([]models.Mission, error) {
	rows, err := q.db.QueryContext(ctx, listMissionsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms []models.Mission
	var prevM models.Mission

	for rows.Next() {
		var t models.Target
		var m models.Mission
		var catID sql.NullInt32
		var notes sql.NullString

		if err := rows.Scan(
			&m.ID,
			&catID,
			&m.IsCompleted,
			&t.ID,
			&t.Name,
			&t.Country,
			&notes,
			&t.IsCompleted,
		); err != nil {
			return nil, err
		}
		m.CatID = catID.Int32
		t.Notes = notes.String

		if m.ID != prevM.ID {
			if prevM.ID != 0 {
				ms = append(ms, prevM)
			}
			prevM = m
		}

		prevM.Targets = append(prevM.Targets, t)
	}
	if prevM.ID != 0 {
		ms = append(ms, prevM)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

const getMissionByCatIdQuery = `
SELECT id, cat_id, is_completed FROM missions
WHERE cat_id = $1;
`

func (q *storage) GetMissionByCatId(ctx context.Context, id int32) (models.Mission, error) {
	row := q.db.QueryRowContext(ctx, getMissionByCatIdQuery, id)
	var m models.Mission
	var catID sql.NullInt32
	err := row.Scan(
		&m.ID,
		&catID,
		&m.IsCompleted,
	)
	m.CatID = catID.Int32
	return m, err
}

const updateMissionCatByIdQuery = `
UPDATE missions 
SET cat_id = $2 
WHERE id = $1
RETURNING id, cat_id, is_completed
`

func (q *storage) UpdateMissionCatById(ctx context.Context, id int32, catID int32) (models.Mission, error) {
	sqlCatID := newNullInt32(catID)
	row := q.db.QueryRowContext(ctx, updateMissionCatByIdQuery, id, sqlCatID)
	var m models.Mission
	err := row.Scan(&m.ID, &sqlCatID, &m.IsCompleted)
	m.CatID = sqlCatID.Int32
	return m, err
}

const updateMissionCompletionByIdQuery = `
UPDATE missions 
SET is_completed = $2 
WHERE id = $1
RETURNING id, cat_id, is_completed
`

func (q *storage) UpdateMissionCompletionById(ctx context.Context, id int32, isCompleted bool) (models.Mission, error) {
	var catID sql.NullInt32
	row := q.db.QueryRowContext(ctx, updateMissionCompletionByIdQuery, id, isCompleted)
	var m models.Mission
	err := row.Scan(&m.ID, &catID, &m.IsCompleted)
	m.CatID = catID.Int32
	return m, err
}
