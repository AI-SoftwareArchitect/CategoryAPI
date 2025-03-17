package routes

import (
    "food-api/internal/interfaces/http/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(router *gin.Engine, controller *controllers.CategoryController) {
    categoryRoutes := router.Group("/api/categories") // /api/categories prefix
    {
        categoryRoutes.POST("", controller.Create)       // POST /api/categories
        categoryRoutes.GET("", controller.GetAll)        // GET /api/categories
        categoryRoutes.GET("/:id", controller.GetByID)   // GET /api/categories/:id
        categoryRoutes.PUT("/:id", controller.Update)    // PUT /api/categories/:id
        categoryRoutes.DELETE("/:id", controller.Delete) // DELETE /api/categories/:id
    }
}