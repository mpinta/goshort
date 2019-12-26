package data

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"goshort/backend/config"
)

func Open() (*gorm.DB, error) {
	c := config.GetConfig()
	return gorm.Open(c.Database.Type, c.Database.Path)
}

func Recreate() error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	db.DropTableIfExists(&Url{})
	db.CreateTable(&Url{})
	return nil
}

func Insert(url Url) error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	db.NewRecord(url)
	db.Create(&url)
	return nil
}

func Get(url string) ([]Url, error) {
	var urls []Url

	db, err := Open()
	if err != nil {
		return urls, err
	}
	defer db.Close()

	rows, err := db.Model(&Url{}).Where("short_url = ?", url).Select("*").Rows()
	if err != nil {
		return urls, err
	}

	urls, err = GetUrlsFromRows(db, rows)
	if err != nil {
		return urls, err
	}

	return urls, nil
}

func GetUrlsFromRows(db *gorm.DB, rows *sql.Rows) ([]Url, error) {
	var urls []Url

	for rows.Next() {
		var oUrl Url
		err := db.ScanRows(rows, &oUrl)
		if err != nil {
			return urls, err
		}

		urls = append(urls, oUrl)
	}

	return urls, nil
}