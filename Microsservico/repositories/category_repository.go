package repositories

import (
	"github.com/Faracoeng/GoLang/ms-categories/Internal/entities"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

// Aqui estamos garantindo que CategoryRepository implementa a interface ICategoryRepository
var _ ICategoryRepository = &CategoryRepository{}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Save(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindById(id uint) (*entities.Category, error) {
	var category entities.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) FindByName(name string) (*entities.Category, error) {
	var category entities.Category
	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) FindAll() ([]*entities.Category, error) {
	var categories []*entities.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) Update(category *entities.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Category{}, id).Error
}
