-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,              -- Product name
    price NUMERIC(10, 2) NOT NULL,           -- Price of the product with precision
    created_at TIMESTAMPTZ DEFAULT now()     -- Timestamp of product creation
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
