package util

import (
	"log"
	"os"
)

// Port gets env port value.
func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8082"
	}
	return ":" + port
}

// MustGetenv gets env variable.
func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}
