package domain

type Movie struct {
	Id         int
	Title      string
	Rating     int
	Details    string
	CategoryID int
}

type MovieWithCategory struct {
	Movie
	CategoryName    string
	CategoryDetails string
}
