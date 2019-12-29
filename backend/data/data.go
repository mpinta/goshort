package data

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"goshort/backend/config"
)

const trigger = "CREATE TRIGGER delete_trigger BEFORE INSERT ON urls FOR EACH ROW BEGIN DELETE FROM urls " +
	"WHERE valid_until < (SELECT strftime('%Y-%m-%d %H:%M:%S', datetime('now', 'localtime'))); END;"

func Open() (*gorm.DB, error) {
	cfg := config.GetConfig()
	return gorm.Open(cfg.Database.Type, cfg.Database.Path)
}

func Recreate() error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	db.DropTableIfExists(&Url{})
	db.CreateTable(&Url{})
	db.Exec(trigger)
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
		var url Url
		err := db.ScanRows(rows, &url)
		if err != nil {
			return urls, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}
