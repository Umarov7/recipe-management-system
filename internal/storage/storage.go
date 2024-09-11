package storage

import (
	"context"
	"recipe-management/internal/models"
)

type IStorage interface {
	Recipe() IRecipeStorage
	Close()
}

type IRecipeStorage interface {
	Create(ctx context.Context, recipe *models.NewRecipe) (string, error)
	Read(ctx context.Context, id string) (*models.Recipe, error)
	Update(ctx context.Context, recipe *models.NewRecipeData) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, p *models.Pagination) ([]*models.RecipeInfo, error)
}
