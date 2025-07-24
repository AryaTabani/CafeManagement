package repository

import (
	"context"
	"database/sql"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func CreateMenuItem(ctx context.Context, item *models.MenuItem) (int64, error) {
	query := `INSERT INTO menu_items (name, description, price, image_url, discount, category_id, is_active) VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := db.DB.ExecContext(ctx, query, item.Name, item.Description, item.Price, item.Image_url, item.Discount, item.Category_id, true)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func UpdateMenuItem(ctx context.Context, item *models.MenuItem) error {
	query := `UPDATE menu_items SET name = ?, description = ?, price = ?, image_url = ?, discount = ?, category_id = ?, is_active = ? WHERE id = ?`
	_, err := db.DB.ExecContext(ctx, query, item.Name, item.Description, item.Price, item.Image_url, item.Discount, item.Category_id, item.ID)
	return err
}

func DeleteMenuItem(ctx context.Context, id int64) error {
	query := `UPDATE menu_items SET is_active = false WHERE id = ?`
	_, err := db.DB.ExecContext(ctx, query, id)
	return err
}
func GetMenuItemByID(ctx context.Context, id int64) error {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM menu_items WHERE id = ?)"
	err := db.DB.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return sql.ErrNoRows
	}
	return nil
}
