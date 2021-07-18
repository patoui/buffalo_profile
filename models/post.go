package models

import (
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

// Post is used by pop to map your posts database table to your go code.
type Post struct {
	ID          int            `json:"id" db:"id"`
	Title       string         `json:"title" db:"title"`
	Slug        string         `json:"slug" db:"slug"`
	Body        string         `json:"body" db:"body"`
	PublishedAt mysql.NullTime `json:"published_at" db:"published_at"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Post) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Posts is not required by pop and may be deleted
type Posts []Post

// Get a post short body (first 100 characters)
func (p Post) ShortBody() string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	sb := strings.TrimSpace(reg.ReplaceAllString(p.Body, ""))
	bl := len(sb) - 1
	if bl > 99 {
		bl = 99
	}
	return sb[0:bl]
}

// Get pretty formatted date time
func (p Post) ShortPublishedAt() string {
	return p.PublishedAt.Time.Format("Mon, Jan _2, 2006 3:04 PM")
}

// Get iso8601 (RFC3339) date time
func (p Post) IsoPublishedAt() string {
	return p.PublishedAt.Time.Format(time.RFC3339)
}

// String is not required by pop and may be deleted
func (p Posts) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Post) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Title, Name: "Title"},
		&validators.StringLengthInRange{Field: p.Title, Name: "Title", Min: 2, Max: 255, Message: "Must be a minimum of 2 characters and maximum of 255"},
		&validators.StringIsPresent{Field: p.Slug, Name: "Slug"},
		&validators.StringLengthInRange{Field: p.Slug, Name: "Slug", Min: 2, Max: 255, Message: "Must be a minimum of 2 characters and maximum of 255"},
		&validators.StringIsPresent{Field: p.Body, Name: "Body"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Post) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Post) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
