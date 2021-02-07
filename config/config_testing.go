// +build testing

package config

import "fmt"

var (
	API_PORT    = "8080"
	DB_DRIVER   = "postgres"
	DB_USER     = "deezefy-music"
	DB_PASSWORD = "deezefy-music"
	DB_DATABASE = "deezefy-music"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "5432"
	DB_SSL      = "disable"
	DB_URL      = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		DB_DRIVER,
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_DATABASE,
		DB_SSL)
)
