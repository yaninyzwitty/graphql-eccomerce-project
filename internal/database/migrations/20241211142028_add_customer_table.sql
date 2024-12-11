-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,              -- Name of the customer
    email VARCHAR(255) UNIQUE NOT NULL,      -- Unique email
    created_at TIMESTAMPTZ DEFAULT now()     -- Timestamp of customer creation
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
