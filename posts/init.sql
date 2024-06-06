CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    post_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    post_text TEXT NOT NULL
);