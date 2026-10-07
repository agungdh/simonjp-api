package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Item struct {
	bun.BaseModel `bun:"table:items"`

	ID          int64     `bun:",pk,autoincrement" json:"id" example:"1"`
	Name        string    `bun:",notnull" json:"name" example:"buku"`
	Description string    `bun:",nullzero" json:"description" example:"belajar go"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

// ItemInput is the request body for create/update.
type ItemInput struct {
	Name        string `json:"name" example:"buku"`
	Description string `json:"description" example:"belajar go"`
}

// ErrorResponse is the standard error payload.
type ErrorResponse struct {
	Error string `json:"error" example:"item not found"`
}
