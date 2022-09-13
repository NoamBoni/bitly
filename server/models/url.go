package models

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/NoamBoni/bitly/server/helpers"
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

func (u *Url) Update() error {
	helpers.LoadEnv()
	u.Modified_url = strings.Replace(u.Modified_url, os.Getenv("DOMAIN"), "", 1)
	if strings.Contains(u.Modified_url, "/") {
		return errors.New("invalid URL")
	}
	switch {
	case u.Modified_url == "" && u.Original_url == "":
		fmt.Println("0")
		return nil
	case u.Modified_url == "" && u.Original_url != "":
		_, err := Db.Model(u).Column("original_url").WherePK().Update()
		fmt.Println("1")
		return err
	case u.Modified_url != "" && u.Original_url == "":
		_, err := Db.Model(u).Column("modified_url").WherePK().Update()
		fmt.Println("2")
		return err
	default:
		_, err := Db.Model(u).Column("modified_url").Column("original_url").WherePK().Update()
		fmt.Println("3")
		return err
	}
}

func GetUrlsByUserId(id int) ([]*Url, error) {
	var urls []*Url
	if err := Db.Model(&urls).Order("created_at ASC").Where("user_id = ?", id).Select(); err != nil {
		return nil, err
	}
	AddDomain(urls)
	return urls, nil
}

func AddDomain(urls []*Url) {
	helpers.LoadEnv()
	for _, url := range urls {
		url.Modified_url = os.Getenv("DOMAIN") + url.Modified_url
	}
}
