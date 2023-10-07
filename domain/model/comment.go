package model

import (
	"time"
)

type Comment struct {
	ID        uint64 `gorm:"primarykey"`
	Article   Article
	ArticleID uint64
	User      User
	UserID    uint64
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
