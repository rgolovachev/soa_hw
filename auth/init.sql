CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password_hash BYTEA NOT NULL
);

CREATE TABLE IF NOT EXISTS user_sessions (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    token TEXT NOT NULL,
    created TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_data (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    firstname TEXT,
    lastname TEXT,
    birthdate TEXT,
    email TEXT,
    telephone TEXT
)