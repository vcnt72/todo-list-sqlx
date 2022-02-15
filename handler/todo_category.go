package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vcnt72/todo-list-sqlx/db/model"
	"github.com/vcnt72/todo-list-sqlx/request"
)

func FindTodoCategory(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		todoCategories := []model.TodoCategory{}

		sql := "SELECT * FROM todo_categories"

		if err := db.Select(&todoCategories, sql); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err,
			})
			return
		}

		var todoCategoryIds []string

		for _, v := range todoCategories {
			todoCategoryIds = append(todoCategoryIds, v.ID)
		}

		query, args, err := sqlx.In("SELECT * FROM todos WHERE todo_category_id IN (?)", todoCategoryIds)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
			return
		}

		query = sqlx.Rebind(sqlx.DOLLAR, query)

		var todos []model.Todo

		if err := db.Select(&todos, query, args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err.Error(),
			})
			return
		}

		for i := range todoCategories {

			for _, v2 := range todos {
				if v2.TodoCategoryID == todoCategories[i].ID {
					todoCategories[i].Todos = append(todoCategories[i].Todos, v2)
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "Success",
			"data":       todoCategories,
		})

	}
}

func CreateTodoCategory(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createDTO request.CreateTodoCategoryDTO
		if err := c.ShouldBindJSON(&createDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"statusCode": http.StatusBadRequest,
				"message":    err,
			})
			return
		}

		id, _ := uuid.NewV4()

		todoCategory := model.TodoCategory{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ID:        id.String(),
			Name:      createDTO.Name,
		}

		sql := "INSERT INTO todo_categories(id, created_at, updated_at, name) VALUES($1,$2,$3,$4)"

		if _, err := db.Exec(sql, todoCategory.ID, todoCategory.CreatedAt, todoCategory.UpdatedAt, todoCategory.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    err,
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"statusCode": http.StatusCreated,
			"message":    "Success",
		})
	}
}

func NewTodoCategoryGroup(base *gin.RouterGroup, db *sqlx.DB) {
	todoCategoryRoute := base.Group("/todo-categories")
	{
		todoCategoryRoute.POST("", CreateTodoCategory(db))
		todoCategoryRoute.GET("", FindTodoCategory(db))
	}
}
