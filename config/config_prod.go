// +build prod

package config

import "os"

var (
	API_PORT  = os.Getenv("PORT")
	DB_DRIVER = os.Getenv("DATABASE_DRIVER")
	DB_HEROKU = os.Getenv("DATABASE_URL")
	DB_URL    = DB_HEROKU + "?sslmode=require"
)
