package endpoints

import (
	"bytes"
	"gofast/api/mocks"
	"io/ioutil"
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

	testClient := GetMockClient(
		t,
		"http://artii.herokuapp.com/make?text=Gotta%20Go%20Fast!",
		200,
		`{"status": 200, message": "Gotta go fast!"}`,
	)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		HealthCheck(w,r, *testClient)
	})

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

	testClient := GetMockClient(
		t,
		"http://artii.herokuapp.com/make?text=Gotta%20Go%20Fast!",
		200,
		`{"status": 200, message": "Gotta go fast!"}`,
	)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		Ping(w,r, *testClient)
	})

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

	testClient := GetMockClient(
		t,
		"http://artii.herokuapp.com/make?text=Gotta%20Go%20Fast!",
		200,
		`{"status": 200, message": "Gotta go fast!"}`,
	)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		TeaPot(w,r, *testClient)
	})

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

	testClient := GetMockClient(
		t,
		"http://artii.herokuapp.com/make?text=Gotta%20Go%20Fast!",
		200,
		`Gotta go fast!`,
		)

	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		GoFast(w,r, *testClient)
	})

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("endpoint returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":200, message":"Gotta go fast!"}`
	if rr.Body.String() != expected {
		t.Errorf("endpoint returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Function to create MockClient to
func GetMockClient(t *testing.T, expectedUrl string, expectedStatusCode int, expectedResponse string) *http.Client {
	testClient := mocks.MockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != expectedUrl {
			t.Fatal("Url expected: " + expectedUrl + ", given: " + req.URL.String())
		}
		return &http.Response{
			StatusCode: expectedStatusCode,
			// Send response to be tested
			Body:       ioutil.NopCloser(bytes.NewBufferString(expectedResponse)),
			// Must be set to non-nil value or it panics
			Header:     make(http.Header),
		}
	})

	return testClient
}
