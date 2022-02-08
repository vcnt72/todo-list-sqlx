CREATE TABLE IF NOT EXISTS todo_categories (
    id uuid PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    name varchar
)