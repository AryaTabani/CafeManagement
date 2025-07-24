package models

type MenuItem struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Final_price float64 `json:"final_price"`
	Image_url   string  `json:"image_url"`
	Discount    int     `json:"discount"` 
	Category_id int64   `json:"-"`
	Is_active   bool    `json:"is_active"`
}

type MenuCategory struct {
	Category_name string     `json:"category_name"`
	Items         []MenuItem `json:"items"`
}