package models

import "fmt"

type Url struct {
	Id           int      `json:"id"`
	Original_url string   `json:"original_url" pg:",notnull"`
	Modified_url string   `json:"modified_url" pg:",notnull"`
	User_id      string   `json:"user_id" pg:",notnull"`
}

func (u *Url) String() string {
	return fmt.Sprintf("%+v\n", *u)
}

func (u *Url) Insert() error {
	_, err := Db.Model(u).Returning("*").Insert()
	return err
}
