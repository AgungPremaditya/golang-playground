package models

type Movie struct {
	MovieID     int64  `json:"movie_id"`
	Title       string `json:"title"`
	Rating      int    `json:"rating"`
	Description string `json:"desc"`
	CategoryID  int64  `json:"category_id"`
}

type MovieWithCategory struct {
	Movie
	CategoryName    string `json:"category_name"`
	CategoryDetails string `json:"category_details"`
}
