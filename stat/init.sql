CREATE TABLE IF NOT EXISTS likes (
    post_id String,
    author_id UInt64,
    user_id UInt64
) ENGINE = MergeTree()
ORDER BY post_id;

CREATE TABLE IF NOT EXISTS views (
    post_id String,
    author_id UInt64,
    user_id UInt64
) ENGINE = MergeTree()
ORDER BY post_id;
