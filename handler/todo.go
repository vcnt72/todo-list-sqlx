package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vcnt72/todo-list-sqlx/db/model"
	"github.com/vcnt72/todo-list-sqlx/request"
)

func FindTodo(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var todos []model.Todo

		if err := db.Select(&todos, `SELECT todos.*,
									 todo_categories.id as "todo_categories.id",
									 todo_categories.name as "todo_categories.name",
									 todo_categories.created_at as "todo_categories.created_at",
									 todo_categories.updated_at as "todo_categories.updated_at"
									 FROM todos INNER JOIN todo_categories ON todos.todo_category_id = todo_categories.id`); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    "Unknown error",
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"statusCode": http.StatusCreated,
			"message":    "Success",
			"data":       todos,
		})
	}
}

func CreateTodo(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createDTO request.CreateTodoDTO
		if err := c.ShouldBindJSON(&createDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"statusCode": http.StatusBadRequest,
				"message":    err.Error(),
			})
			return
		}

		todoCategory := model.TodoCategory{}

		findSql := "SELECT * FROM todo_categories WHERE id = $1"

		if err := db.Get(&todoCategory, findSql, createDTO.TodoCategoryID); err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"statusCode": http.StatusNotFound,
					"message":    "Todo category not found",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
			return
		}

		id, _ := uuid.NewV4()

		todo := model.Todo{
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			ID:             id.String(),
			Title:          createDTO.Title,
			Description:    createDTO.Description,
			TodoCategoryID: createDTO.TodoCategoryID,
		}

		sql := "INSERT INTO todos(id, created_at, updated_at, title, description, todo_category_id) VALUES($1,$2,$3,$4,$5,$6)"

		if _, err := db.Exec(sql, todo.ID, todo.CreatedAt, todo.UpdatedAt, todo.Title, todo.Description, todo.TodoCategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"statusCode": http.StatusCreated,
			"message":    "Success",
		})
	}
}

func UpdateTodo(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var updateDTO request.UpdateTodoDTO
		if err := c.ShouldBindJSON(&updateDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"statusCode": http.StatusBadRequest,
				"message":    err.Error(),
			})
			return
		}

		todoCategory := model.TodoCategory{}

		findTodoCategorySql := "SELECT id FROM todo_categories WHERE id = $1"

		if err := db.Get(&todoCategory, findTodoCategorySql, updateDTO.TodoCategoryID); err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"statusCode": http.StatusNotFound,
					"message":    "Todo category not found",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
			return
		}

		findTodoSql := "SELECT id FROM todos WHERE id = $1"

		todo := model.Todo{}

		if err := db.Get(&todo, findTodoSql, id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"statusCode": http.StatusNotFound,
					"message":    "Todo category not found",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})

			return
		}

		sql := "UPDATE todos SET updated_at=$1, title=$2, description=$3, todo_category_id=$4 WHERE id=$5"

		if _, err := db.Exec(sql, time.Now(), updateDTO.Title, updateDTO.Description, updateDTO.TodoCategoryID, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "Success",
		})
	}
}

func NewTodoRoute(baseGroup *gin.RouterGroup, db *sqlx.DB) {
	todoGroup := baseGroup.Group("/todos")
	{
		todoGroup.POST("", CreateTodo(db))
		todoGroup.PUT("/:id", UpdateTodo(db))
		todoGroup.GET("", FindTodo(db))
	}
}
