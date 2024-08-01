CREATE TABLE rates (
                       id SERIAL PRIMARY KEY,
                       created_at TIMESTAMP WITH TIME ZONE,
                       updated_at TIMESTAMP WITH TIME ZONE,
                       deleted_at TIMESTAMP WITH TIME ZONE,
                       currency_from VARCHAR(255) NOT NULL,
                       currency_to VARCHAR(255) NOT NULL,
                       rate FLOAT8 NOT NULL
);
