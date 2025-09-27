# Cafe Management REST API

This project is a REST API designed to manage a cafe's menu and administrative functions. It is built with **Go** and the **Gin** framework, providing a secure and efficient backend for an admin panel.

## ✨ Features

-   **Secure Authentication**: User login is protected with **JWT** (JSON Web Tokens) and **Bcrypt** for secure password hashing.
-   **Role-Based Access**: The API distinguishes between public routes (e.g., viewing the menu) and protected admin-only routes for management tasks.
-   **Full Menu Management (CRUD)**: Admins have full Create, Read, Update, and Delete capabilities for menu items and their categories.
-   **Soft Deletes**: Deleting a menu item is a "soft delete," meaning the item is marked as inactive rather than being permanently removed from the database, allowing for data recovery.
-   **Layered Architecture**: The project follows a clean 3-layer architecture (Controller, Service, Repository) to ensure separation of concerns and maintainable code.

## 🛠️ Tech Stack

-   **Language**: Go
-   **Web Framework**: Gin
-   **Database**: SQLite
-   **Authentication**: JWT & Bcrypt

## 📄 API Endpoints

### Public Routes

| Method | Path  | Description               | Auth Required |
| :----- | :---- | :------------------------ | :-----------: |
| `GET`  | `/menu` | Fetches the public menu |      No       |

### Admin Routes (Protected)

| Method   | Path                    | Description                      | Auth Required |
| :------- | :---------------------- | :------------------------------- | :-----------: |
| `POST`   | `/admin/login`          | Admin login to get an auth token |      No       |
| `POST`   | `/admin/menu-items`     | Create a new menu item           |      Yes      |
| `PUT`    | `/admin/menu-items/:id` | Update an existing menu item     |      Yes      |
| `DELETE` | `/admin/menu-items/:id` | Deactivate (soft delete) an item |      Yes      |
| `POST`   | `/admin/categories`     | Create a new menu category       |      Yes      |
| `PUT`    | `/admin/categories/:id` | Update a menu category           |      Yes      |
| `DELETE` | `/admin/categories/:id` | Delete a menu category           |      Yes      |

## 🚀 Getting Started

**Prerequisites:**
-   Go (version 1.18 or later)
-   Git

**Instructions:**
1.  Clone the repository:
    ```bash
    git clone https://github.com/AryaTabani/CafeManagement.git
    cd CafeManagement
    ```
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Run the application:
    ```bash
    go run main.go
    ```
    The service will be running at `http://localhost:8080`.
