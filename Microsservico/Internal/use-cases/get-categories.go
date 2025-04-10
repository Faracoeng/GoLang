package use_cases

import (
	"github.com/Faracoeng/GoLang/ms-categories/Internal/entities"
	"github.com/Faracoeng/GoLang/ms-categories/repositories"
)

type GetCategoriesUseCase struct {
	repository repositories.ICategoryRepository
}

func NewGetCategoriesUseCase(repository repositories.ICategoryRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{repository}
}

func (u *GetCategoriesUseCase) Execute() ([]*entities.Category, error) {
	return u.repository.FindAll()
}
