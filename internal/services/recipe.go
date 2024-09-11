package services

import (
	"log/slog"
	"recipe-management/internal/storage"
)

type RecipeService struct {
	storage *storage.IRecipeStorage
	logger  *slog.Logger
}

func NewRecipeService(storage *storage.IRecipeStorage, logger *slog.Logger) *RecipeService {
	return &RecipeService{
		storage: storage,
		logger:  logger,
	}
}
