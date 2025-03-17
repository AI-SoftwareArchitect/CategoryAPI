package repositories

import (
	"food-api/internal/domain/entities"
	"food-api/pkg/utils"
)

type CategoryRepository interface {
    Create(category *entities.Category) error
    Update(category *entities.Category) error
    Delete(id string) error
    GetByID(id string) (*entities.Category, error)
    GetAll(pagination *utils.Pagination) ([]entities.Category, error)
}
