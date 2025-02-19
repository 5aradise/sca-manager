CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    years_of_experience INT NOT NULL,
    breed VARCHAR(255) NOT NULL,
    salary_in_cents INT NOT NULL
);

CREATE TABLE missions (
    id SERIAL PRIMARY KEY,
    cat_id INT,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (cat_id) REFERENCES cats(id) ON DELETE SET NULL
);

CREATE TABLE targets (
    id SERIAL PRIMARY KEY,
    mission_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    notes TEXT,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (mission_id) REFERENCES missions(id) ON DELETE CASCADE
);
