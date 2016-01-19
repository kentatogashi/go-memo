package models

import (
	"time"
)

func init() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = loc
}

type Content struct {
	Id        int64
	Content   string `sql:"type:text(2000)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
