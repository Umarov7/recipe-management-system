package storage

import (
	"context"
	"recipe-management/internal/models"
)

type IStorage interface {
	User() IUserStorage
	Recipe() IRecipeStorage
	Close()
}

type IUserStorage interface {
	Create(ctx context.Context, user *models.NewUser) (string, error)
	Read(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.NewUserData) error
	Delete(ctx context.Context, id string) error
}

type IRecipeStorage interface {
	Create(ctx context.Context, recipe *models.NewRecipe) (string, error)
	Read(ctx context.Context, id string) (*models.Recipe, error)
	Update(ctx context.Context, recipe *models.NewRecipeData) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, p *models.RecipeFilter) ([]*models.RecipeInfo, error)
}
