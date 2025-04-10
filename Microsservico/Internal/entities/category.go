package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID   uint    `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category :=  &Category {
		Name: name,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	// Aqui então devem ser adicionados os métodos para manipular a categoria, como adicionar, editar e remover.
	// e definir as regras de negócio
	err := category.idValid()
	if err != nil {	
		return nil, err

	}
	return category, nil
}


func (c *Category) idValid() error {
	if  len(c.Name) < 5 {
		return fmt.Errorf("nome da categoria deve pelo menos ser maior que 5 caracteres, está com: %d", len(c.Name))
	}
	return nil
}

func (c *Category) UpdateName(newName string) (*Category, error) {
	if len(newName) < 5 {
		return nil, fmt.Errorf("nome da categoria deve ter pelo menos 5 caracteres, recebeu: %d", len(newName))
	}
	c.Name = newName
	c.UpdatedAt = time.Now().Format(time.RFC3339)
	return c, nil
}
