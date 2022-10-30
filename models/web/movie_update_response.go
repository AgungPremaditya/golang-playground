package web

type MovieUpdateRequest struct {
	Id         int
	Title      string `validate:"required"`
	Rating     int    `validate:"required"`
	Details    string
	CategoryId int `validate:"required"`
}
