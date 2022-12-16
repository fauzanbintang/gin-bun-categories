package services

import (
	"context"
	"database/sql"
	"zamannow/go-rest-api/db"
	"zamannow/go-rest-api/domain"
	"zamannow/go-rest-api/domain/repository"
	"zamannow/go-rest-api/dto/requests"
	"zamannow/go-rest-api/dto/responses"

	"github.com/uptrace/bun"
)

type CategoryService interface {
	Create(ctx context.Context, src requests.CreateCategoryRequest) (res domain.Category, err error)
	Update(ctx context.Context, src requests.UpdateCategoryRequest) (res responses.CreateCategoryResponse, err error)
	Delete(ctx context.Context, id int64) (err error)
	FindById(ctx context.Context, id int64) (res domain.Category, err error)
	FindAll(ctx context.Context) (res []domain.Category, err error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (srv *categoryService) Create(ctx context.Context, src requests.CreateCategoryRequest) (res domain.Category, err error) {
	c, cancel := repository.NewContext(ctx)
	defer cancel()

	category := domain.Category{
		Name: src.Name,
	}

	if err = db.GetConn().RunInTx(c, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) (err error) {
		if err = srv.categoryRepo.Create(ctx, &tx, &category); err != nil {
			panic(err)
		}

		return nil
	}); err != nil {
		panic(err)
	}

	return category, nil
}

func (srv *categoryService) Update(ctx context.Context, src requests.UpdateCategoryRequest) (res responses.CreateCategoryResponse, err error) {
	c, cancel := repository.NewContext(ctx)
	defer cancel()

	category := domain.Category{
		ID:   src.ID,
		Name: src.Name,
	}

	if err = db.GetConn().RunInTx(c, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) (err error) {
		if err = srv.categoryRepo.Update(ctx, &tx, &category); err != nil {
			panic(err)
		}

		return nil
	}); err != nil {
		panic(err)
	}

	res.ID = category.ID
	res.Name = category.Name

	return res, nil
}

func (srv *categoryService) Delete(ctx context.Context, id int64) (err error) {
	c, cancel := repository.NewContext(ctx)
	defer cancel()

	category := domain.Category{
		ID: id,
	}

	if err = db.GetConn().RunInTx(c, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) (err error) {
		if err = srv.categoryRepo.Delete(ctx, &tx, &category); err != nil {
			panic(err)
		}

		return
	}); err != nil {
		panic(err)
	}

	return
}

func (srv *categoryService) FindById(ctx context.Context, id int64) (res domain.Category, err error) {
	c, cancel := repository.NewContext(ctx)
	defer cancel()

	res.ID = id

	if err = db.GetConn().RunInTx(c, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) (err error) {
		if err = srv.categoryRepo.FindById(ctx, &tx, &res); err != nil {
			panic(err)
		}

		return
	}); err != nil {
		panic(err)
	}

	return
}

func (srv *categoryService) FindAll(ctx context.Context) (res []domain.Category, err error) {
	c, cancel := repository.NewContext(ctx)
	defer cancel()

	if err = db.GetConn().RunInTx(c, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) (err error) {
		if err = srv.categoryRepo.FindAll(ctx, &tx, &res); err != nil {
			panic(err)
		}

		return
	}); err != nil {
		panic(err)
	}

	return
}
