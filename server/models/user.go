package models

import (
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname" pg:",notnull" binding:"required"`
	Lastname  string `json:"lastname" pg:",notnull" binding:"required"`
	Email     string `pg:",unique,notnull" json:"email" binding:"required"`
	Password  string `json:"password" pg:",notnull" binding:"required"`
	Urls      []*Url `pg:"-,rel:has-many" json:"urls"`
}

func (u *User) String() string {
	return fmt.Sprintf("%+v", *u)
}

func (u *User) Insert() error {
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return err
	}
	if hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		u.Password = string(hash)
	}
	if _, err := Db.Model(u).Returning("id, firstname, lastname, email").Insert(); err != nil {
		return err
	}
	u.Password = ""
	return nil
}

func (u *User) SelectByEmail() error {
	if err := Db.Model(u).Column("*").Where("email = ?",u.Email).Select(); err != nil {
		return err
	}
	return nil
}

func (u *User) SelectById() error {
	if err := Db.Model(u).Column("*").Where("id = ?",u.Id).Select(); err != nil {
		return err
	}
	return nil
}
