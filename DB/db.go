package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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
	createDefaultAdmin()
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
		discount INTEGER, 
		is_active BOOLEAN NOT NULL DEFAULT true,
		category_id INTEGER,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);`

	_, err = DB.Exec(createMenuItemsTable)
	if err != nil {
		panic("Failed to create menu_items table")
	}
	createAdminsTable := `
    CREATE TABLE IF NOT EXISTS admins (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL
    );`

	_, err = DB.Exec(createAdminsTable)
	if err != nil {
		panic("Failed to create admins table")
	}

}
func createDefaultAdmin() {
	defaultUsername := "admin"
	defaultPassword := "admin123"

	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM admins WHERE username = ?", defaultUsername).Scan(&count)
	if err != nil {
		panic("Failed to check for default admin: " + err.Error())
	}
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
		if err != nil {
			panic("Failed to hash default password: " + err.Error())
		}
		query := `INSERT INTO admins (username, password_hash) VALUES (?, ?)`
		_, err = DB.Exec(query, defaultUsername, string(hashedPassword))
		if err != nil {
			panic("Failed to create default admin user: " + err.Error())
		}
		log.Printf("   Username: %s\n", defaultUsername)
		log.Printf("   Password: %s\n", defaultPassword)
	}
}
