package repository

import (
	"context"
	"time"
)

type repositoryPool struct {
	Category CategoryRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		Category: NewCategoryRepository(),
	}
}

func NewContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, 5*time.Second)
}
