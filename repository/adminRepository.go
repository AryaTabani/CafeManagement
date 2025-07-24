package repository

import (
	"context"
	"database/sql"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func GetAdminByUsername(ctx context.Context, username string) (*models.Admin, error) {
	var admin models.Admin
	query := "SELECT id, username, password_hash FROM admins WHERE username = ?"
	err := db.DB.QueryRowContext(ctx, query, username).Scan(&admin.ID, &admin.Username, &admin.Password_hash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}
	return &admin, nil
}