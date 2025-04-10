package repositories

import "github.com/Faracoeng/GoLang/ms-categories/Internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	FindById(id uint) (*entities.Category, error)
	FindByName(name string) (*entities.Category, error)
	FindAll() ([]*entities.Category, error)
	Update(category *entities.Category) error
	Delete(id uint) error
}
