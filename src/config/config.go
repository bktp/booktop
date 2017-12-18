package config

import (
	"fmt"
	"net/http"
)

const (
	dbUser     = "buktp"
	dbPassword = "password"
	dbName     = "booktop"
)

// SendError replies with status text
func SendError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

// ConnectionString returns psql connection string
func ConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
}
