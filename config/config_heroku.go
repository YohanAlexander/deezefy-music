// +build heroku

package config

import (
	"fmt"
	"os"
)

var (
	API_PORT  = os.Getenv("PORT")
	DB_DRIVER = os.Getenv("DATABASE_DRIVER")
	DB_HEROKU = os.Getenv("DATABASE_URL")
	DB_SSL    = os.Getenv("DATABASE_SSL")
	DB_URL    = DB_HEROKU + fmt.Sprintf("?sslmode=%s", DB_SSL)
)
