package repository

import (
	"context"
	"zamannow/go-rest-api/db"
	"zamannow/go-rest-api/domain"

	"github.com/uptrace/bun"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error)
	Update(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error)
	Delete(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error)
	FindById(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error)
	FindAll(ctx context.Context, tx *bun.Tx, src *[]domain.Category) (err error)
}

type categoryRepository struct {
	db *bun.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		db: db.GetConn(),
	}
}

func (r *categoryRepository) Create(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error) {
	_, err = tx.NewInsert().
		Model(src).
		Returning("*").
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *categoryRepository) Update(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error) {
	_, err = tx.NewUpdate().
		Model(src).
		Where("id = ?", src.ID).
		Returning("*").
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error) {
	_, err = tx.NewDelete().
		Model(src).
		Where("id = ?", src.ID).
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *categoryRepository) FindById(ctx context.Context, tx *bun.Tx, src *domain.Category) (err error) {
	err = tx.NewSelect().
		Model(src).
		Where("id = ?", src.ID).
		Scan(ctx)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *categoryRepository) FindAll(ctx context.Context, tx *bun.Tx, src *[]domain.Category) (err error) {
	err = tx.NewSelect().
		Model(src).
		Scan(ctx)
	if err != nil {
		panic(err)
	}

	return nil
}
