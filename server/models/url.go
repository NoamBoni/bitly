package models

import (
	"fmt"
	"net/url"
	"time"
	"errors"
)

type Url struct {
	Id           int       `json:"id"`
	Original_url string    `json:"original_url" pg:",notnull" binding:"required"`
	Modified_url string    `json:"modified_url" pg:",unique,notnull"`
	User_id      int       `json:"user_id" pg:",notnull fk:user_id on_delete:CASCADE"`
	Created_at   time.Time `pg:"default:now()" json:"created_at"`
}

func (u *Url) String() string {
	return fmt.Sprintf("%+v", *u)
}

func (u *Url) Insert() error {
	parsedUrl, err := url.ParseRequestURI(u.Original_url)
	if err != nil {
		return err
	}
	if parsedUrl.Host == "" || parsedUrl.Scheme == "" {
		return errors.New("invalid URL")
	}
	_, err = Db.Model(u).Returning("*").Insert()
	return err
}

func GetUrlsByUserId(id int) ([]Url, error) {
	var urls []Url
	if err := Db.Model(&urls).Order("created_at ASC").Where("user_id = ?", id).Select(); err != nil {
		return nil, err
	}
	return urls, nil
}
