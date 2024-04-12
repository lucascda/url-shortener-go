CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    hash TEXT NOT NULL,    
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id INT REFERENCES users(id)
);