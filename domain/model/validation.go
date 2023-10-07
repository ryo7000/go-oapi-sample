package model

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"

	validation "github.com/go-ozzo/ozzo-validation"
)

func checkUnique(tx *gorm.DB, failMessage string) validation.RuleFunc {
	return func(value interface{}) error {
		var count int64
		result := tx.Count(&count)
		if result.Error != nil {
			slog.Warn("Count error", "error", result.Error.Error())
			return errors.New("unknown error")
		}

		if count > 0 {
			return errors.New(failMessage)
		}

		return nil
	}
}
