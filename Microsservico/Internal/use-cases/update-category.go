package use_cases

import (
	"errors"
	"github.com/Faracoeng/GoLang/ms-categories/repositories"
)

type UpdateCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewUpdateCategoryUseCase(repository repositories.ICategoryRepository) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{repository}
}

type UpdateCategoryInput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (u *UpdateCategoryUseCase) Execute(input UpdateCategoryInput) error {
	category, err := u.repository.FindById(input.ID)
	if err != nil {
		return err
	}

	if category == nil {
		return errors.New("category not found")
	}

	if category.Name == input.Name {
		return nil
	}

	conflict, err := u.repository.FindByName(input.Name)
	if err != nil {
		return err
	}
	if conflict != nil && conflict.ID != input.ID {
		return errors.New("name already in use")
	}

	updated, err := category.UpdateName(input.Name)
	if err != nil {
		return err
	}

	return u.repository.Update(updated)
}
