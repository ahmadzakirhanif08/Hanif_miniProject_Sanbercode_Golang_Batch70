-- Up migration
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INTEGER NOT NULL REFERENCES categories (id) ON DELETE RESTRICT,
    description TEXT,
    image_url VARCHAR(255),
    release_year INTEGER CHECK (release_year BETWEEN 1980 AND 2024),
    price INTEGER NOT NULL,
    total_page INTEGER NOT NULL,
    thickness VARCHAR(50),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    created_by VARCHAR(100),
    modified_at TIMESTAMP WITHOUT TIME ZONE,
    modified_by VARCHAR(100)
);

-- Down migration
DROP TABLE IF EXISTS books;