package usecases

import (
	"food-api/internal/domain/entities"
	"food-api/internal/domain/repositories"
	"food-api/pkg/utils"
)

type CategoryUsecase struct {
    repo repositories.CategoryRepository
}

func NewCategoryUsecase(repo repositories.CategoryRepository) *CategoryUsecase {
    return &CategoryUsecase{repo: repo}
}

func (uc *CategoryUsecase) CreateCategory(category *entities.Category) error {
    return uc.repo.Create(category)
}

func (uc *CategoryUsecase) GetCategoryByID(id string) (*entities.Category, error) {
    return uc.repo.GetByID(id)
}

func (uc *CategoryUsecase) GetAllCategories(pagination *utils.Pagination) ([]entities.Category, error) {
    return uc.repo.GetAll(pagination)
}

func (uc *CategoryUsecase) UpdateCategory(category *entities.Category) error {
    return uc.repo.Update(category)
}

func (uc *CategoryUsecase) DeleteCategory(id string) error {
    return uc.repo.Delete(id)
}