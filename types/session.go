package types

type Session struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Email  string `json:"email"`
}
