CREATE TABLE IF NOT EXISTS item (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    unit varchar(255) NOT NULL,
    amount double precision NOT NULL,
    expires_at timestamptz NOT NULL,
);
