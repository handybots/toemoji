CREATE DATABASE IF NOT EXISTS toemoji; -- use your name

CREATE TABLE IF NOT EXISTS toemoji.logs (
    date        Date,
    time        DateTime,
    level       String,
    message     String,
    event       String,
    user_id     UInt32
) ENGINE = MergeTree(date, (level, event, user_id), 8192);
