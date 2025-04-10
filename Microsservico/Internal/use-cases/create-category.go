package use_cases

import (
	"github.com/Faracoeng/GoLang/ms-categories/Internal/entities"
	"github.com/Faracoeng/GoLang/ms-categories/repositories"
)

type CreateCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{repository}
}

func (u *CreateCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	return u.repository.Save(category)
}
