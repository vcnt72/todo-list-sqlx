package db

import (
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func NewConnection() *sqlx.DB {
	databaseUrl := viper.GetString("db-url")

	connConfig, _ := pgx.ParseConfig(databaseUrl)

	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sqlx.Open("pgx", connStr)

	if err != nil {
		fmt.Println("Unable to connect database")
		panic(err)
	}

	return db
}
