package models

type AdminCategoryPayload struct {
	Name      string `json:"name" binding:"required"`
	Icon_url  string `json:"icon_url"`
}