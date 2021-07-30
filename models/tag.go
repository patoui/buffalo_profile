package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/pop/v5/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

// Tag is used by pop to map your tags database table to your go code.
type Tag struct {
	ID          int        `json:"id" db:"id"`
	Name        slices.Map `json:"name" db:"name"`
	Slug        slices.Map `json:"slug" db:"slug"`
	Type        string     `json:"type" db:"type"`
	OrderColumn int        `json:"order_column" db:"order_column"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Tag) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tags is not required by pop and may be deleted
type Tags []Tag

// String is not required by pop and may be deleted
func (t Tags) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Tag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: t.ID, Name: "ID"},
		&validators.StringIsPresent{Field: t.Type, Name: "Type"},
		&validators.IntIsPresent{Field: t.OrderColumn, Name: "OrderColumn"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Tag) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Tag) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
