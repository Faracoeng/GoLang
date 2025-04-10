package use_cases

import (
	"github.com/Faracoeng/GoLang/ms-categories/repositories"
)

type DeleteCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewDeleteCategoryUseCase(repository repositories.ICategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{repository}
}

func (u *DeleteCategoryUseCase) Execute(id uint) error {
	return u.repository.Delete(id)
}
