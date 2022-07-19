#!/bin/bash
export PG_PRESENCE_URL="postgresql://postgres:password@localhost:5432/presences"
export JWT_PRESENCE_SECRET="YourJWTSecret"
go run  ./main.go