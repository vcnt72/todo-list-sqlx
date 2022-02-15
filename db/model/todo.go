package model

import "time"

type Todo struct {
	ID             string    `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	TodoCategoryID string    `json:"todo_category_id" db:"todo_category_id"`
	TodoCategory   `json:"todo_category" db:"todo_categories"`
}
