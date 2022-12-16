package requests

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
