package request

type CreateTodoDTO struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	TodoCategoryID string `json:"todo_category_id"`
}

type UpdateTodoDTO struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	TodoCategoryID string `json:"todo_category_id"`
}
