package main

import (
	"gofast/api/endpoints"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", func (w http.ResponseWriter, r *http.Request) { endpoints.HealthCheck(w,r) })
	http.HandleFunc("/ping", func (w http.ResponseWriter, r *http.Request) { endpoints.Ping(w,r) })
	http.HandleFunc("/teapot", func (w http.ResponseWriter, r *http.Request) { endpoints.TeaPot(w,r) })
	http.HandleFunc("/api/gofast", func (w http.ResponseWriter, r *http.Request) { endpoints.GoFast(w,r) })
	http.ListenAndServe(":3003", nil)
}
