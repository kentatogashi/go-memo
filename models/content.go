package models

import (
	"time"
)

type Content struct {
	Id        int64
	Content   string `sql:"type:text(2000)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
