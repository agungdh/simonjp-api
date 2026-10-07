package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Item struct {
	bun.BaseModel `bun:"table:items"`

	ID          int64     `bun:",pk,autoincrement" json:"id"`
	Name        string    `bun:",notnull" json:"name"`
	Description string    `bun:",nullzero" json:"description"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
