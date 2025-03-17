package services

import (
	"food-api/internal/domain/entities"
	"food-api/internal/domain/repositories"
	"food-api/pkg/utils"
)

type CategoryService struct {
    repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) *CategoryService {
    return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *entities.Category) error {
    return s.repo.Create(category)
}

func (s *CategoryService) GetCategoryByID(id string) (*entities.Category, error) {
    return s.repo.GetByID(id)
}

func (s *CategoryService) GetAllCategories(pagination *utils.Pagination) ([]entities.Category, error) {
    return s.repo.GetAll(pagination)
}

func (s *CategoryService) UpdateCategory(category *entities.Category) error {
    return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id string) error {
    return s.repo.Delete(id)
}