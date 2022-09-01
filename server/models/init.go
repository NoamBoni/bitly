package models

import (
	"context"
	"os"

	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var Db *pg.DB

func init() {
	helpers.LoadEnv()
	Db = pg.Connect(&pg.Options{
		Addr:      os.Getenv("DB_URL"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASS"),
		Database:  os.Getenv("DB_NAME"),
		TLSConfig: nil,
	})
	ctx := context.Background()
	if _, err := Db.ExecContext(ctx, "SELECT 1"); err != nil {
		panic(err)
	}
	err := createSchema(Db)
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Url)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
