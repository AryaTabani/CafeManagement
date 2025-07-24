package repository

import (
	"context"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error) {
	query := `
		SELECT
			mi.id, mi.name, mi.description, mi.price, mi.image_url, mi.discount, mi.category_id
		FROM menu_items mi
		WHERE mi.is_active = true -- Only fetch active items
		ORDER BY mi.category_id, mi.id
	`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.MenuItem
	for rows.Next() {
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Image_url, &item.Discount, &item.Category_id)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func GetCategoryNames(ctx context.Context) (map[int64]string, error) {
	query := "SELECT id, name FROM categories"
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categoryNames := make(map[int64]string)
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		categoryNames[id] = name
	}
	return categoryNames, nil
}