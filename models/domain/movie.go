package domain

type Movie struct {
	Id         int
	Title      string
	Rating     int
	Details    string
	CategoryId int
}

type MovieWithCategory struct {
	Id              int
	Title           string
	Rating          int
	Details         string
	CategoryId      int
	CategoryName    string
	CategoryDetails string
}
