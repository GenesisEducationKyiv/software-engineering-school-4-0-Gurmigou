CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       created_at TIMESTAMP WITH TIME ZONE,
                       updated_at TIMESTAMP WITH TIME ZONE,
                       deleted_at TIMESTAMP WITH TIME ZONE,
                       email VARCHAR(255) UNIQUE NOT NULL
);
