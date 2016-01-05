package models

import (
	"time"
)

type Content struct {
	Id        int64
	Content   string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
