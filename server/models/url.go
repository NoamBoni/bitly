package models

import (
	"fmt"
	"time"
)

type Url struct {
	Id           int       `json:"id"`
	Original_url string    `json:"original_url" pg:",notnull" binding:"required"`
	Modified_url string    `json:"modified_url" pg:",notnull" binding:"required"`
	User_id      string    `json:"user_id" pg:",notnull"`
	CreatedAt    time.Time `pg:"default:now()"`
}

func (u *Url) String() string {
	return fmt.Sprintf("%+v", *u)
}

func (u *Url) Insert() error {
	_, err := Db.Model(u).Returning("*").Insert()
	return err
}
