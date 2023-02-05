CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS inventory (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(255) NOT NULL,
    unit varchar(255) NOT NULL,
    amount double precision NOT NULL,
    expires_at timestamptz NOT NULL
);
