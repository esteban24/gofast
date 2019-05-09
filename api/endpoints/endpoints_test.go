package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("endpoint returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status": 200, "alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("endpoint returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPing(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("endpoint returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status": 200, "message": "pong"}`
	if rr.Body.String() != expected {
		t.Errorf("endpoint returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTeaPot(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/teapot", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TeaPot)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusTeapot {
		t.Errorf("endpoint returned wrong status code: got %v want %v",
			status, http.StatusTeapot)
	}

	// Check the response body is what we expect.
	expected := `{"status": 418, message": "I'm a teapot"}`
	if rr.Body.String() != expected {
		t.Errorf("endpoint returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGoFast(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/teapot", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GoFast)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("endpoint returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status": 200, message": "Gotta go fast!"}`
	if rr.Body.String() != expected {
		t.Errorf("endpoint returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
