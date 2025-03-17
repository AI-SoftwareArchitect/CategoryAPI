package unit

import (
    "food-api/internal/application/services"
    "food-api/internal/domain/entities"
    "food-api/internal/domain/repositories"
    "testing"
)

type MockCategoryRepo struct {
    repositories.CategoryRepository
}

func (m *MockCategoryRepo) Exists(name string) (bool, error) {
    return name == "existing", nil
}

func TestCreateCategory(t *testing.T) {
    mockRepo := &MockCategoryRepo{}
    service := services.NewCategoryService(mockRepo)

    t.Run("Create new category", func(t *testing.T) {
        category := &entities.Category{Name: "test"}
        err := service.CreateCategory(category)
        if err != nil {
            t.Errorf("Unexpected error: %v", err)
        }
    })
}