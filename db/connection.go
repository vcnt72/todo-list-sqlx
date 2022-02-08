package db

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func NewConnection() *sqlx.DB {
	databaseUrl := viper.GetString("db-url")

	db, err := sqlx.Open("pgx", databaseUrl)

	if err != nil {
		fmt.Println("Unable to connect database")
		panic(err)
	}

	return db
}
