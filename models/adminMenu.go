package models

type AdminMenuItemPayload struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Discount    int    `json:"discount"`
	Final_price float64 `json:"final_price"`
	Image_url   string `json:"image_url"`
	Category_id int64  `json:"category_id" binding:"required"`
	Is_active   bool   `json:"is_active"`
}
