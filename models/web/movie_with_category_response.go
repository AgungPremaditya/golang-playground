package web

type MovieWithCategoryResponse struct {
	Id              int
	Title           string
	Rating          int
	Details         string
	CategoryId      int
	CategoryName    string
	CategoryDetails string
}
