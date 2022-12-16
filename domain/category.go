package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	Name      string    `bun:"name" json:"name"`
	CreatedAt time.Time `bun:"created_at" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at" json:"updated_at"`
}

func (m *Category) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.UpdatedAt = time.Now()
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
