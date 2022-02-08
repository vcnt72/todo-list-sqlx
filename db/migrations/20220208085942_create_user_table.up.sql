CREATE TABLE IF NOT EXISTS users  (
    id uuid PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    email varchar,
    name varchar
);