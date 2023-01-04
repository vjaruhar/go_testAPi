package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetAllThreads(t *testing.T) {
	req, err := http.NewRequest("GET", "/get_all_threads", nil)
	if err != nil {
		//t.Fatal(err)
		t.Log(req)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllThreads)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		t.Log(rr.Body.String())
	} else {
		//fmt.Printf("handler returned correct status code: got %v want %v and body %v", status, http.StatusOK, rr)
		t.Log(rr.Body.String())
	}
}

func TestCreateThread(t *testing.T) {

	data := url.Values{}
	data.Set("name", "1")
	data.Set("description", "1")

	req, err := http.NewRequest("POST", "/create_thread", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateThread)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	t.Log(rr.Body.String())
}

func TestGetThread(t *testing.T) {

	//var jsonStr = []byte(`{"threadID":"1"}`)

	data := url.Values{}
	data.Set("threadID", "1")

	req, err := http.NewRequest("POST", "/get_thread", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetThread)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	t.Logf("%v", rr.Body.String())
}
