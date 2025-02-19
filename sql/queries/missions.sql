INSERT INTO missions (cat_id, is_completed) 
VALUES ($1, $2)
RETURNING id, cat_id, is_completed;

SELECT m.id, m.cat_id, m.is_completed,
    t.id AS target_id, t.name, t.country, t.notes, t.is_completed
FROM missions m
LEFT JOIN targets t ON m.id = t.mission_id
WHERE m.id = $1;

SELECT m.id, m.cat_id, m.is_completed,
       t.id AS target_id, t.name, t.country, t.notes, t.is_completed
FROM missions m
LEFT JOIN targets t ON m.id = t.mission_id
ORDER BY m.id;

SELECT * FROM missions
WHERE cat_id = $1;

UPDATE missions 
SET cat_id = $2 
WHERE id = $1
RETURNING id, cat_id, is_completed;

UPDATE missions 
SET is_completed = $2 
WHERE id = $1
RETURNING id, cat_id, is_completed;

DELETE FROM missions WHERE id = $1;
