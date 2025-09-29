-- Up migration
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    created_by VARCHAR(100),
    modified_at TIMESTAMP WITHOUT TIME ZONE,
    modified_by VARCHAR(100)
);

-- Down migration
DROP TABLE IF EXISTS users;