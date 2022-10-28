package models

type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Details      string `json:"details"`
}
