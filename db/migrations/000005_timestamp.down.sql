ALTER TABLE categories
ALTER COLUMN created_at TYPE timestamp,
ALTER COLUMN updated_at TYPE timestamp;

ALTER TABLE groups
ALTER COLUMN created_at TYPE timestamp,
ALTER COLUMN updated_at TYPE timestamp;

ALTER TABLE movies
ALTER COLUMN created_at TYPE timestamp,
ALTER COLUMN updated_at TYPE timestamp;

ALTER TABLE themes
ALTER COLUMN created_at TYPE timestamp,
ALTER COLUMN updated_at TYPE timestamp;

ALTER TABLE tracks
ALTER COLUMN created_at TYPE timestamp,
ALTER COLUMN updated_at TYPE timestamp;

ALTER TABLE users
ALTER COLUMN created_at TYPE timestamp,
ALTER COLUMN updated_at TYPE timestamp;
