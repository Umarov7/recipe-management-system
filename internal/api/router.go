package api

import (
	_ "recipe-management/internal/api/docs"
	"log/slog"
	"recipe-management/internal/api/handler"
	"recipe-management/internal/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Recipe Management System
// @version 1.0
// @description API for Recipe Management System
// @host localhost:8080
// @BasePath /recipes
// @schemes http
func NewRouter(storage storage.IStorage, logger *slog.Logger) *gin.Engine {
	h := handler.NewHandler(storage, logger)
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/recipes", h.Add)
	router.GET("/recipes/:id", h.Get)
	router.PUT("/recipes/:id", h.Update)
	router.DELETE("/recipes/:id", h.Delete)
	router.GET("/recipes", h.List)

	return router
}
