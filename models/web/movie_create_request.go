package web

type MovieCreateRequest struct {
	Title      string
	Rating     int
	Details    string
	CategoryId int
}
