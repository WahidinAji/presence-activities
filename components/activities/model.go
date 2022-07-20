package activities

import "github.com/jackc/pgx/v4"

type ActivityIn struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ActivityOut struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

//findall Activity
type ListIn struct {
	UserId int `json:"user_id"`
}

type ActivitiesList struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

//activity deps
type ActivityDeps struct {
	DB *pgx.Conn
}
