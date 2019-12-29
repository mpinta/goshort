package data

import "time"

type Url struct {
	Id           int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	FullUrl      string    `gorm:"type:VARCHAR(100);NOT NULL"`
	ShortUrl     string    `gorm:"type:VARCHAR(50);NOT NULL"`
	CreatedAt    time.Time `gorm:"type:DATETIME;NOT NULL"`
	ValidUntil   time.Time `gorm:"type:DATETIME;NOT NULL"`
	MinutesValid int       `gorm:"NOT NULL"`
}

func (Url) TableName() string {
	return "urls"
}
