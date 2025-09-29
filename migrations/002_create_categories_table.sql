-- Up migration
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    created_by VARCHAR(100),
    modified_at TIMESTAMP WITHOUT TIME ZONE,
    modified_by VARCHAR(100)
);

-- Down migration
DROP TABLE IF EXISTS categories;