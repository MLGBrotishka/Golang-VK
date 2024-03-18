CREATE TABLE IF NOT EXISTS actors (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gender ENUM('male', 'female', 'unknown') NOT NULL,
    birth_date DATE NOT NULL
    movies TEXT
);