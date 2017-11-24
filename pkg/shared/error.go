package shared

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var (
	// ErrNoResult is a not results error
	ErrNoResult = errors.New("Result not found.")
	// ErrUnavailable is a database not available error
	ErrUnavailable = errors.New("Database is unavailable.")
	// ErrUnauthorized is a permissions violation
	ErrUnauthorized = errors.New("User does not have permission to perform this operation.")
)

// standardizeErrors returns the same error regardless of the database used
func StandardizeError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return ErrNoResult
	}
	return err
}