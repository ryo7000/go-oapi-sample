package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

// For model validation
var _ callbacks.BeforeSaveInterface = (*User)(nil)

type Article struct {
	ID          uint64 `gorm:"primarykey"`
	Author      User   `gorm:"foreignKey:AuthorID"`
	AuthorID    uint64
	Tags        []Tag `gorm:"many2many:article_tags"`
	Slug        string
	Title       string
	Description string
	Body        string
	Favorites   []User `gorm:"many2many:usr_favorites"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tag struct {
	Description string
	Articles    []Article `gorm:"many2many:article_tags"`
	CreatedAt   time.Time
}

func (a *Article) BeforeSave(db *gorm.DB) error {
	if err := a.Validate(db); err != nil {
		return err
	}

	return nil
}

func (a *Article) Validate(db *gorm.DB) error {
	return validation.ValidateStruct(a,
		validation.Field(
			&a.AuthorID,
			validation.Required.Error("ユーザが選択されていません"),
		),
	)
}
