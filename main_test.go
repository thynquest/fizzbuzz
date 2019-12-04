package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz/{multint1:[0-9]+}/{multint2:[0-9]+}/{limit:[0-9]+}/{multstr1:[a-z]+}/{multstr2:[a-z]+}", fizzbuzzHandler).Methods("GET")
	router.HandleFunc("/stats", statsHandler).Methods("GET")
	return router
}

func TestFizzbuzzHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := Router()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/10/fizz/buzz", nil))
	if w.Code != http.StatusOK {
		t.Errorf("wrong HTTP status code: expected %d; got %d", 200, w.Code)
	}
	body := w.Body.String()
	if len(body) == 0 {
		t.Error("Expected not empty string")
	}
	expected := "[\"1\",\"2\",\"fizz\",\"4\",\"buzz\",\"fizz\",\"7\",\"8\",\"fizz\",\"buzz\"]"
	if strings.TrimSuffix(body, "\n") != expected {
		t.Errorf("Wrong string received, expected %s, got %s", expected, body)
	}
}

func TestFizzbuzzWrongNumberParams(t *testing.T) {
	w := httptest.NewRecorder()
	r := Router()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/10/buzz", nil))
	if w.Code != http.StatusNotFound {
		t.Errorf("wrong HTTP status code: expected %d; got %d", 200, w.Code)
	}
}

func TestFizzbuzzWrongTypeParams(t *testing.T) {
	w := httptest.NewRecorder()
	r := Router()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/aa/10/fizz/buzz", nil))
	if w.Code != http.StatusNotFound {
		t.Errorf("wrong HTTP status code: expected %d; got %d", 404, w.Code)
	}
}

func TestFizzbuzzNegativeLimit(t *testing.T) {
	w := httptest.NewRecorder()
	r := Router()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/-1/fizz/buzz", nil))
	if w.Code != http.StatusNotFound {
		t.Errorf("wrong HTTP status code: expected %d; got %d", 404, w.Code)
	}
}

func TestFizzbuzzZeroMultiple(t *testing.T) {
	w := httptest.NewRecorder()
	r := Router()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/0/0/100/fizz/buzz", nil))
	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong HTTP status code: expected %d; got %d", 404, w.Code)
	}
}

func TestStatHandler(t *testing.T) {
	w := httptest.NewRecorder()
	z := httptest.NewRecorder()
	r := Router()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/100/fizz/buzz", nil))
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/100/fizz/buzz", nil))
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/100/fizz/buzz", nil))
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/200/fizz/buzz", nil))
	r.ServeHTTP(w, httptest.NewRequest("GET", "/fizzbuzz/3/5/200/fizz/buzz", nil))
	r.ServeHTTP(z, httptest.NewRequest("GET", "/stats", nil))
	if z.Code != http.StatusOK {
		t.Errorf("wrong HTTP status code: expected %d; got %d", 200, w.Code)
	}
	body := z.Body.String()
	expected := "{\"Query\":\"3,5,100,fizz,buzz\",\"Hits\":3}"
	if len(body) == 0 {
		t.Error("Expected not empty string")
	}
	if strings.TrimSuffix(body, "\n") != expected {
		t.Errorf("Wrong string received, expected %s, got %s", expected, body)
	}
}

func TestStatsWithNoFizzBuzzRequest(t *testing.T) {
	w := httptest.NewRecorder()
	r := Router()
	for k := range frequency {
		delete(frequency, k)
	}
	r.ServeHTTP(w, httptest.NewRequest("GET", "/stats", nil))
	if w.Code != http.StatusOK {
		t.Errorf("wrong HTTP status code: expected %d; got %d", 200, w.Code)
	}
	body := w.Body.String()
	expected := "no fizzbuzz request registered"
	if len(body) == 0 {
		t.Error("Expected not empty string")
	}
	if body != expected {
		t.Errorf("Wrong message received expected %s got %s", body, expected)
	}
}
