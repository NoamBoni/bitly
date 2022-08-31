package models

import (
	"context"
	"os"

	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/go-pg/pg/v10"
)

var Db *pg.DB

func init() {
	Db = ConnectDB()
}

func ConnectDB() *pg.DB {
	helpers.LoadEnv()
	db := pg.Connect(&pg.Options{
		Addr:      os.Getenv("DB_URL"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASS"),
		Database:  os.Getenv("DB_NAME"),
		TLSConfig: nil,
	})
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, "SELECT 1"); err != nil {
		panic(err)
	}
	return db
}
