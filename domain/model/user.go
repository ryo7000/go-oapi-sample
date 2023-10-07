package model

import (
	"errors"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

// For model validation
var _ callbacks.BeforeSaveInterface = (*User)(nil)

type User struct {
	ID         uint64 `gorm:"primarykey"`
	Username   string
	Email      string
	Password   string
	Bio        *string
	Image      *string
	Followers  []Follow  `gorm:"foreignKey:FollowingID"`
	Followings []Follow  `gorm:"foreignKey:FollowerID"`
	Favorites  []Article `gorm:"many2many:user_favorites"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Follow struct {
	Follower    User
	FollowerID  uint64
	Following   User
	FollowingID uint64
}

func (u *User) BeforeSave(db *gorm.DB) error {
	_, err := bcrypt.Cost([]byte(u.Password))
	if err != nil {
		panic("Password not hashed. Please use HashPassword()")
	}

	if err := u.Validate(db); err != nil {
		return err
	}

	return nil
}

func (u *User) Validate(db *gorm.DB) error {
	return validation.ValidateStruct(u,
		validation.Field(
			&u.Username,
			validation.Required.Error("名前の入力が必要です"),
			validation.Match(regexp.MustCompile("[a-zA-Z0-9]+")).Error("小文字と大文字の英数字のみ使えます"),
			validation.By(
				checkUnique(
					db.Model(&User{}).
						Where(&User{Username: u.Username}),
					"入力された名前はすでに使用されています",
				),
			),
		),
		validation.Field(
			&u.Email,
			validation.Required.Error("メールアドレスの入力が必要です"),
			is.Email.Error("メールアドレスを入力して下さい"),
			validation.By(
				checkUnique(
					db.Model(&User{}).
						Where(&User{Email: u.Email}),
					"入力されたメールアドレスはすでに使用されています",
				),
			),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("パスワードの入力が必要です"),
		),
	)
}

func HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

func (u *User) FollowedBy(id uint64) bool {
	if u.Followers == nil {
		return false
	}
	for _, f := range u.Followers {
		if f.FollowerID == id {
			return true
		}
	}
	return false
}
