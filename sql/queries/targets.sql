INSERT INTO targets (mission_id, name, country, notes, is_completed) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, country, notes, is_completed;

UPDATE targets
SET is_completed = $2 
WHERE id = $1
RETURNING id, name, country, notes, is_completed;

UPDATE targets
SET notes = $2 
WHERE id = $1
RETURNING id, mission_id, name, country, notes, is_completed;

DELETE FROM targets WHERE id = $1;
