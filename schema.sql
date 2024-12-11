CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    email STRING UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    price FLOAT  NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()

);

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASACADE,
    created_at TIMESTAMPTZ DEFAULT now()
);
CREATE TABLE IF NOT EXISTS order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASACADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASACADE,
    quantity INT NOT NULL,
    price FLOAT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);


-- for postgres

CREATE EXTENSION IF NOT EXISTS pgcrypto;  -- Ensure pgcrypto extension is available for gen_random_uuid()

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,              -- Name of the customer
    email VARCHAR(255) UNIQUE NOT NULL,      -- Unique email
    created_at TIMESTAMPTZ DEFAULT now()     -- Timestamp of customer creation
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,              -- Product name
    price NUMERIC(10, 2) NOT NULL,           -- Price of the product with precision
    created_at TIMESTAMPTZ DEFAULT now()     -- Timestamp of product creation
);

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now()     -- Timestamp of order creation
);

CREATE TABLE IF NOT EXISTS order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),  -- Ensure positive quantity
    price NUMERIC(10, 2) NOT NULL,               -- Price at the time of the order
    created_at TIMESTAMPTZ DEFAULT now()         -- Timestamp of order item creation
);
