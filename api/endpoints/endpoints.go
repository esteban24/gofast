package endpoints

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type ResponseMessage struct {
	Status 	int		`json:"status"`
	Message string	`json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request, client http.Client) {
	var errors []error

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, writeErr := io.WriteString(w, `{"status": 200, "alive": true}`)

	errors = append(errors, writeErr)
	HandleError(errors)
}

func Ping(w http.ResponseWriter, r *http.Request, client http.Client) {
	var errors []error

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, writeErr := io.WriteString(w, `{"status": 200, "message": "pong"}`)

	errors = append(errors, writeErr)
	HandleError(errors)
}

func TeaPot(w http.ResponseWriter, r *http.Request, client http.Client) {
	var errors []error

	w.WriteHeader(http.StatusTeapot)
	w.Header().Set("Content-Type", "application/json")

	_, writeErr := io.WriteString(w, `{"status": 418, message": "I'm a teapot"}`)

	errors = append(errors, writeErr)
	HandleError(errors)
}

func GoFast(w http.ResponseWriter, r *http.Request, client http.Client) {
	var errors []error

	request, reqErr := http.NewRequest("GET", "http://artii.herokuapp.com/make?text=Gotta%20Go%20Fast!", nil)

	response, clientErr := client.Do(request)
	body, readErr := ioutil.ReadAll(response.Body)


	w.WriteHeader(response.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	var payload = `{"status":` + strconv.Itoa(response.StatusCode) + `, message":"` + string(body)  + `"}`

	_, writeErr := io.WriteString(w, payload)

	errors = append(errors, readErr, reqErr, clientErr, writeErr)
	HandleError(errors)
}

func HandleError(errors []error) {
	for _, err := range errors {
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Normalize(val []byte) ResponseMessage {
	var message ResponseMessage
	s, _ := strconv.Unquote(string(val))
	json.Unmarshal([]byte(s), &message)

	return message
}
