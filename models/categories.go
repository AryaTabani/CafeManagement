package models

type Category struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Icon_url string `json:"icon_url"`
}