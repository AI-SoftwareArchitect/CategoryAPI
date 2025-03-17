package cassandra

import (
	"food-api/internal/domain/entities"
	"food-api/internal/domain/repositories"
	"food-api/pkg/utils"

	"github.com/gocql/gocql"
)

type categoryRepository struct {
    session *gocql.Session
}

func NewCategoryRepository(session *gocql.Session) repositories.CategoryRepository {
    return &categoryRepository{session: session}
}

func (r *categoryRepository) Create(category *entities.Category) error {
    return r.session.Query(
        `INSERT INTO category (categoryid, categoryname, categorydescription) VALUES (?, ?, ?)`,
        category.ID,
        category.Name,
        category.Description,
    ).Exec()
}

func (r *categoryRepository) GetByID(id string) (*entities.Category, error) {
    var category entities.Category
    err := r.session.Query(
        `SELECT categoryid, categoryname, categorydescription FROM category WHERE categoryid = ?`,
        id,
    ).Scan(&category.ID, &category.Name, &category.Description)
    return &category, err
}

func (r *categoryRepository) GetAll(pagination *utils.Pagination) ([]entities.Category, error) {
    var categories []entities.Category
    iter := r.session.Query(
        `SELECT categoryid, categoryname, categorydescription FROM category LIMIT ?`,
        pagination.Limit,
    ).Iter()

    var category entities.Category
    for iter.Scan(&category.ID, &category.Name, &category.Description) {
        categories = append(categories, category)
    }

    return categories, iter.Close()
}

func (r *categoryRepository) Update(category *entities.Category) error {
    return r.session.Query(
        `UPDATE category SET categoryname = ?, categorydescription = ? WHERE categoryid = ?`,
        category.Name,
        category.Description,
        category.ID,
    ).Exec()
}

func (r *categoryRepository) Delete(id string) error {
    return r.session.Query(
        `DELETE FROM category WHERE categoryid = ?`,
        id,
    ).Exec()
}