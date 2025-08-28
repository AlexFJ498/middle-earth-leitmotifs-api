DROP TRIGGER IF EXISTS update_categories_timestamps ON categories;
DROP TRIGGER IF EXISTS update_groups_timestamps ON groups;
DROP TRIGGER IF EXISTS update_movies_timestamps ON movies;
DROP TRIGGER IF EXISTS update_themes_timestamps ON themes;
DROP TRIGGER IF EXISTS update_tracks_timestamps ON tracks;
DROP TRIGGER IF EXISTS update_users_timestamps ON users;

DROP FUNCTION IF EXISTS update_timestamps;

ALTER TABLE categories
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE groups
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE movies
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE themes
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE tracks
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE users
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;
