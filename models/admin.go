package models

type Admin struct {
    ID           int64
    Username     string
    Password_hash string
}

type LoginPayload struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}