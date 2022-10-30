package web

type MovieCreateRequest struct {
	Title      string `validate:"required"`
	Rating     int    `validate:"required"`
	Details    string
	CategoryId int `validate:"required"`
}
