package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createtables()
}

func createtables() {
	createCategoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		icon_url TEXT
	);`

	_, err := DB.Exec(createCategoriesTable)
	if err != nil {
		panic("Failed to create categories table")
	}

	createMenuItemsTable := `
	CREATE TABLE IF NOT EXISTS menu_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		image_url TEXT,
		category_id INTEGER,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);`

	_, err = DB.Exec(createMenuItemsTable)
	if err != nil {
		panic("Failed to create menu_items table")
	}
}