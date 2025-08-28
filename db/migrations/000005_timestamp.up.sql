ALTER TABLE categories
ALTER COLUMN created_at SET DATA TYPE timestamp(0),
ALTER COLUMN updated_at SET DATA TYPE timestamp(0);

ALTER TABLE groups
ALTER COLUMN created_at SET DATA TYPE timestamp(0),
ALTER COLUMN updated_at SET DATA TYPE timestamp(0);

ALTER TABLE movies
ALTER COLUMN created_at SET DATA TYPE timestamp(0),
ALTER COLUMN updated_at SET DATA TYPE timestamp(0);

ALTER TABLE themes
ALTER COLUMN created_at SET DATA TYPE timestamp(0),
ALTER COLUMN updated_at SET DATA TYPE timestamp(0);

ALTER TABLE tracks
ALTER COLUMN created_at SET DATA TYPE timestamp(0),
ALTER COLUMN updated_at SET DATA TYPE timestamp(0);

ALTER TABLE users
ALTER COLUMN created_at SET DATA TYPE timestamp(0),
ALTER COLUMN updated_at SET DATA TYPE timestamp(0);
