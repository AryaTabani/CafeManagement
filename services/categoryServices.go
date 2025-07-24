package services

import (
	"context"
	"errors"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)

var (
	ErrCategoryInUse = errors.New("cannot delete category: it is currently in use by one or more menu items")
)

func CreateCategory(ctx context.Context, payload *models.AdminCategoryPayload) (*models.Category, error) {
	newID, err := repository.CreateCategory(ctx, payload)
	if err != nil {
		return nil, err
	}
	return &models.Category{ID: newID, Name: payload.Name, Icon_url: payload.Icon_url}, nil
}

func GetAllCategories(ctx context.Context) ([]models.Category, error) {
	return repository.GetAllCategories(ctx)
}

func UpdateCategory(ctx context.Context, id int64, payload *models.AdminCategoryPayload) error {
	exists, err := repository.CategoryExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrCategoryNotFound
	}
	return repository.UpdateCategory(ctx, id, payload)
}

func DeleteCategory(ctx context.Context, id int64) error {
	exists, err := repository.CategoryExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrCategoryNotFound
	}

	count, err := repository.CountMenuItemsByCategoryID(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrCategoryInUse
	}

	return repository.DeleteCategory(ctx, id)
}
