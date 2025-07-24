package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)

var (
	ErrMenuItemNotFound = errors.New("menu item not found")
	ErrCategoryNotFound = errors.New("category not found")
)

func CreateMenuItem(ctx context.Context, payload *models.AdminMenuItemPayload) (*models.MenuItem, error) {
	categoryExists, err := repository.CategoryExists(ctx, payload.Category_id)
	if err != nil {
		return nil, fmt.Errorf("failed to verify category existence: %w", err)
	}
	if !categoryExists {
		return nil, ErrCategoryNotFound
	}

	itemToCreate := &models.MenuItem{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Image_url:   payload.Image_url,
		Discount:    payload.Discount,
		Category_id: payload.Category_id,
		Is_active:   payload.Is_active,
	}

	newID, err := repository.CreateMenuItem(ctx, itemToCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create menu item in repository: %w", err)
	}

	itemToCreate.ID = newID
	return itemToCreate, nil
}

func UpdateMenuItem(ctx context.Context, id int64, payload *models.AdminMenuItemPayload) error {
	err := repository.GetMenuItemByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrMenuItemNotFound
		}
		return fmt.Errorf("failed to verify menu item existence: %w", err)
	}

	categoryExists, err := repository.CategoryExists(ctx, payload.Category_id)
	if err != nil {
		return fmt.Errorf("failed to verify category existence: %w", err)
	}
	if !categoryExists {
		return ErrCategoryNotFound
	}

	itemToUpdate := &models.MenuItem{
		ID:          id,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Image_url:   payload.Image_url,
		Discount:    payload.Discount,
		Category_id: payload.Category_id,
		Is_active:   payload.Is_active,
		Final_price: payload.Final_price,
	}

	return repository.UpdateMenuItem(ctx, itemToUpdate)
}

func DeleteMenuItem(ctx context.Context, id int64) error {
	err := repository.GetMenuItemByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrMenuItemNotFound
		}
		return fmt.Errorf("failed to verify menu item existence: %w", err)
	}

	return repository.DeleteMenuItem(ctx, id)
}
