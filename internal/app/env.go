package app

import "os"

func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "development"
}

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}
