package web

type MovieUpdateRequest struct {
	Id         int
	Title      string
	Rating     int
	Details    string
	CategoryId int
}
