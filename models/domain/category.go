package domain

type Category struct {
	Id      int    `json:"category_id"`
	Name    string `json:"category_name"`
	Details string `json:"details"`
}
