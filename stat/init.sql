CREATE TABLE IF NOT EXISTS likes (
    post_id String,
    username String
) ENGINE = MergeTree()
ORDER BY post_id;

CREATE TABLE IF NOT EXISTS views (
    post_id String,
    username String
) ENGINE = MergeTree()
ORDER BY post_id;
