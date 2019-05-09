package endpoints

import (
	"io"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, err := io.WriteString(w, `{"status": 200, "alive": true}`)

	HandleError(err)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, err := io.WriteString(w, `{"status": 200, "message": "pong"}`)

	HandleError(err)
}

func TeaPot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	w.Header().Set("Content-Type", "application/json")

	_, err := io.WriteString(w, `{"status": 418, message": "I'm a teapot"}`)

	HandleError(err)
}

func GoFast(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, err := io.WriteString(w, `{"status": 200, message": "Gotta go fast!"}`)

	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
