package models

import "fmt"

type User struct {
	Id        int      `json:"id"`
	Firstname string   `json:"firstname" pg:",notnull"`
	Lastname  string   `json:"lastname" pg:",notnull"`
	Email     string   `pg:",unique,notnull" json:"email"`
	Password  string   `json:"password" pg:",notnull"`
	Urls      []*Url   `sql:"-" json:"urls"`
}

func (u *User) String() string {
	return fmt.Sprintf("%+v\n", *u)
}

func (u *User) Insert() error {
	_, err := Db.Model(u).Returning("id, firstname, lastname, email").Insert()
	return err
}
