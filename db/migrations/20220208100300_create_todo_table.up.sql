CREATE TABLE IF NOT EXISTS  todos (
    id uuid PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    title varchar,
    description text,
    user_id uuid,
    todo_category_id uuid,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES todo_list.public.users(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_todo_category_id
        FOREIGN KEY (todo_category_id)
            REFERENCES todo_categories(id)
            ON DELETE CASCADE
)