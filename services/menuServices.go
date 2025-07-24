package services

import (
	"context"
	"fmt"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)


func GetFullMenu(ctx context.Context) ([]models.MenuCategory, error) {
	items, err := repository.GetAllMenuItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve menu items: %w", err)
	}

	categoryNames, err := repository.GetCategoryNames(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve category names: %w", err)
	}

	for i := range items {
		item := &items[i]
		if item.Discount > 0 {
			discountAmount := item.Price * (float64(item.Discount) / 100.0)
			item.Final_price = item.Price - discountAmount
		} else {
			item.Final_price = item.Price
		}
	}

	itemsByCategoryID := make(map[int64][]models.MenuItem)
	for _, item := range items {
		itemsByCategoryID[item.Category_id] = append(itemsByCategoryID[item.Category_id], item)
	}
	
	var fullMenu []models.MenuCategory
	for categoryID, categoryItems := range itemsByCategoryID {
		menuCategory := models.MenuCategory{
			Category_name: categoryNames[categoryID],
			Items:         categoryItems,
		}
		fullMenu = append(fullMenu, menuCategory)
	}
	return fullMenu, nil
}