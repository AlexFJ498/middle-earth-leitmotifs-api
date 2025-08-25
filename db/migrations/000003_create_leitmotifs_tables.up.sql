CREATE TABLE movies (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE groups (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE tracks (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    movie_id UUID REFERENCES movies(id) NOT NULL
);

CREATE TABLE themes (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    group_id UUID REFERENCES groups(id) NOT NULL,
    category_id UUID REFERENCES categories(id),
    first_heard UUID REFERENCES tracks(id) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);