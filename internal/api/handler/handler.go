package handler

import (
	"errors"
	"log/slog"
	"recipe-management/internal/storage"
	"strconv"
)

type Handler struct {
	storage storage.IRecipeStorage
	logger  *slog.Logger
}

func NewHandler(s storage.IStorage, l *slog.Logger) *Handler {
	return &Handler{
		storage: s.Recipe(),
		logger:  l,
	}
}

func parseIntQueryParam(queryParam string) (int, error) {
	if queryParam == "" {
		return -1, errors.New("empty integer parameter")
	}

	value, err := strconv.Atoi(queryParam)
	if err != nil || value < 1 {
		return -1, errors.New("invalid integer parameter")
	}

	return value, nil
}
