package controllers

import (
    "food-api/internal/application/services"
    "food-api/internal/domain/entities"
    "food-api/internal/interfaces/http/hateoas"
    "food-api/pkg/utils"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type CategoryController struct {
    service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
    return &CategoryController{service: service}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the input payload
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body entities.Category true "Create category"
// @Success 201 {object} entities.Category
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /categories [post]
func (ctrl *CategoryController) Create(c *gin.Context) {
    var category entities.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.service.CreateCategory(&category); err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }

    c.Header("Location", "/categories/"+strconv.Itoa(category.ID))
    c.JSON(http.StatusCreated, gin.H{
        "data":  category,
        "links": hateoas.GenerateCategoryLinks(c.Request, strconv.Itoa(category.ID)),
    })
}

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description Get a category by its ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} entities.Category
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func (ctrl *CategoryController) GetByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    category, err := ctrl.service.GetCategoryByID(strconv.Itoa(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  category,
        "links": hateoas.GenerateCategoryLinks(c.Request, strconv.Itoa(category.ID)),
    })
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories with pagination
// @Tags categories
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} entities.Category
// @Router /categories [get]
func (ctrl *CategoryController) GetAll(c *gin.Context) {
    var pagination utils.Pagination
    if err := c.ShouldBindQuery(&pagination); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    categories, err := ctrl.service.GetAllCategories(&pagination)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  categories,
        "links": hateoas.GenerateCategoryLinks(c.Request, ""),
    })
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category by its ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body entities.Category true "Update category"
// @Success 200 {object} entities.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [put]
func (ctrl *CategoryController) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    var category entities.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    category.ID = id
    if err := ctrl.service.UpdateCategory(&category); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  category,
        "links": hateoas.GenerateCategoryLinks(c.Request, strconv.Itoa(category.ID)),
    })
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by its ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [delete]
func (ctrl *CategoryController) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    if err := ctrl.service.DeleteCategory(strconv.Itoa(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}