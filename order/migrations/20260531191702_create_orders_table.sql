-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    uuid            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_uuid       UUID NOT NULL,
    part_uuids      UUID[] NOT NULL,
    total_price     DOUBLE PRECISION NOT NULL DEFAULT 0,
    transaction_uuid UUID,
    payment_method  TEXT,
    status          TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ
    );

-- +goose Down
DROP TABLE IF EXISTS orders;
