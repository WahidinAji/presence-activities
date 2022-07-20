package users

import "github.com/jackc/pgx/v4"

//regiser area
type RegisIn struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type RegisOut struct {
	Email string `json:"email"`
}

//login area
type LoginIn struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
type LoginOut struct {
	Email string `json:"email"`
}

//response json area
type JsonResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

//deps area
type UserDeps struct {
	DB *pgx.Conn
}