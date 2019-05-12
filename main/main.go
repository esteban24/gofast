package main

import (
	"gofast/api/endpoints"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health-check", func (w http.ResponseWriter, r *http.Request) {
		endpoints.HealthCheck(w,r, http.Client{Timeout:time.Second * 90})
	})
	http.HandleFunc("/ping", func (w http.ResponseWriter, r *http.Request) {
		endpoints.Ping(w,r, http.Client{Timeout:time.Second * 90})
	})
	http.HandleFunc("/teapot", func (w http.ResponseWriter, r *http.Request) {
		endpoints.TeaPot(w,r, http.Client{Timeout:time.Second * 90})
	})
	http.HandleFunc("/api/gofast", func (w http.ResponseWriter, r *http.Request) {
		endpoints.GoFast(w,r, http.Client{Timeout:time.Second * 90})
	})
	http.ListenAndServe(":3003", nil)
}
