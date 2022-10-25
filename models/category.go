package models

type Category struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	Details      string `json:"details"`
}
