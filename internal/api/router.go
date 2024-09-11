package api

import (
	"log/slog"
	"recipe-management/internal/api/handler"
	"recipe-management/internal/storage"

	"github.com/gin-gonic/gin"
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

	router.POST("/recipes", h.Add)
	router.GET("/recipes/:id", h.Get)
	router.PUT("/recipes/:id", h.Update)
	router.DELETE("/recipes/:id", h.Delete)
	router.GET("/recipes", h.List)

	return router
}
