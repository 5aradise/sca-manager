INSERT INTO cats (name, years_of_experience, breed, salary_in_cents) 
VALUES ($1, $2, $3, $4)
RETURNING id, name, years_of_experience, breed, salary_in_cents;

SELECT id, name, years_of_experience, breed, salary_in_cents FROM cats 
WHERE id = $1;

SELECT id, name, years_of_experience, breed, salary_in_cents FROM cats;

UPDATE cats 
SET salary_in_cents = $2 
WHERE id = $1
RETURNING id, name, years_of_experience, breed, salary_in_cents;

DELETE FROM cats WHERE id = $1;
