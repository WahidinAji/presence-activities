package presences

import (
	"time"

	"github.com/jackc/pgx/v4"
)

type PresenceIn struct {
	// ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Status string `json:"status"`
}

type PresenceOut struct {
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	PresenceCheck time.Time `json:"presence_check"`
}

type ListIn struct{
	UserId int `json:"user_id,string"`
}

type PresenceList struct {
	ID int    `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	PresenceCheck time.Time `json:"presence_check"`
}

type PresenceDeps struct {
	DB *pgx.Conn
}
