package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(g *gin.Engine, db *sqlx.DB) {

	base := g.Group("/api")
	{
		NewTodoRoute(base, db)
		NewTodoCategoryGroup(base, db)
	}

}
