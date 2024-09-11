package services

import (
	"log/slog"
	"recipe-management/internal/storage"
)

type UserService struct {
	storage *storage.IUserStorage
	logger  *slog.Logger
}

func NewUserService(storage *storage.IUserStorage, logger *slog.Logger) *UserService {
	return &UserService{
		storage: storage,
		logger:  logger,
	}
}
