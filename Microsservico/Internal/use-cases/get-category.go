package use_cases

import (
	"errors"

	"github.com/Faracoeng/GoLang/ms-categories/Internal/entities"
	"github.com/Faracoeng/GoLang/ms-categories/repositories"
)

type GetCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewGetCategoryUseCase(repository repositories.ICategoryRepository) *GetCategoryUseCase {
	return &GetCategoryUseCase{repository}
}

func (u *GetCategoryUseCase) Execute(id uint) (*entities.Category, error) {
	category, err := u.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
