package pkg

import "os"

var (
	PG_URL     = os.Getenv("PG_PRESENCE_URL")
	JWTSecret = os.Getenv("JWT_PRESENCE_SECRET")
)

//postgresql://postgres:password@127.0.0.1:5432/presences
