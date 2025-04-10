package controllers

import (
	"net/http"
	"strconv"
	"github.com/Faracoeng/GoLang/ms-categories/repositories"
	"github.com/Faracoeng/GoLang/ms-categories/Internal/use-cases"
	"github.com/gin-gonic/gin"
)

var categoryRepo repositories.ICategoryRepository

func InitCategoryController(repo repositories.ICategoryRepository) {
	categoryRepo = repo
}

func CreateCategory(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase(categoryRepo)
	if err := useCase.Execute(input.Name); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Categoria criada com sucesso"})
}

func GetCategories(c *gin.Context) {
	useCase := use_cases.NewGetCategoriesUseCase(categoryRepo)
	categories, err := useCase.Execute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	useCase := use_cases.NewGetCategoryUseCase(categoryRepo)
	category, err := useCase.Execute(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := use_cases.NewUpdateCategoryUseCase(categoryRepo)
	err = useCase.Execute(use_cases.UpdateCategoryInput{
		ID:   uint(id),
		Name: input.Name,
	})

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoria atualizada com sucesso"})
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	useCase := use_cases.NewDeleteCategoryUseCase(categoryRepo)
	if err := useCase.Execute(uint(id)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoria deletada com sucesso"})
}
