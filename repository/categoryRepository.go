package repository

import (
	"context"
	"database/sql"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func CreateCategory(ctx context.Context, payload *models.AdminCategoryPayload) (int64, error) {
	query := `INSERT INTO categories (name, icon_url) VALUES (?, ?)`
	res, err := db.DB.ExecContext(ctx, query, payload.Name, payload.Icon_url)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := "SELECT id, name, icon_url FROM categories"
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Icon_url); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

func UpdateCategory(ctx context.Context, id int64, payload *models.AdminCategoryPayload) error {
	query := `UPDATE categories SET name = ?, icon_url = ? WHERE id = ?`
	_, err := db.DB.ExecContext(ctx, query, payload.Name, payload.Icon_url, id)
	return err
}

func DeleteCategory(ctx context.Context, id int64) error {
	query := `DELETE FROM categories WHERE id = ?`
	_, err := db.DB.ExecContext(ctx, query, id)
	return err
}

func CategoryExists(ctx context.Context, id int64) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)"
	err := db.DB.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func CountMenuItemsByCategoryID(ctx context.Context, id int64) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM menu_items WHERE category_id = ?"
	err := db.DB.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count, nil
}