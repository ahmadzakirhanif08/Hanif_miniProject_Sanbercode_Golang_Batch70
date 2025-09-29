-- +migrate Up
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INTEGER NOT NULL,
    description VARCHAR(255),
    image_url VARCHAR(255),
    release_year INTEGER,
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255),
    
    -- Foreign Key Constraint
    CONSTRAINT fk_category
        FOREIGN KEY(category_id) 
        REFERENCES categories(id)
        ON DELETE RESTRICT
);

-- +migrate Down
DROP TABLE books;