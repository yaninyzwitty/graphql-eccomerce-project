-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),  -- Ensure positive quantity
    price NUMERIC(10, 2) NOT NULL,               -- Price at the time of the order
    created_at TIMESTAMPTZ DEFAULT now()         -- Timestamp of order item creation
);

CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id); --secondary index improve performance

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
