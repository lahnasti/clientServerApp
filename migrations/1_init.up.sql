CREATE TABLE IF NOT EXISTS users(
    uid serial PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL
);
