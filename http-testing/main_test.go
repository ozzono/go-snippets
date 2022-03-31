package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_CreateProduct(t *testing.T) {
	router, err := run()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/products/create", strings.NewReader(`{"name": "New Product"}`))
	if err != nil {
		t.Log(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		logData(req, t)
	}
}

func Test_ListProducts(t *testing.T) {
	router, err := run()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, "/products/list", nil)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusFound {
		logData(req, t)
	}
}

func Test_FindProduct(t *testing.T) {
	router, err := run()
	if err != nil {
		t.Fatal(err)
	}
	params := url.Values{}
	params.Add("key", "1")
	u := "/products/find?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		t.Log("http.NewRequest", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusFound {
		logData(req, t)
	}
}

func Test_FindUnexistantProduct(t *testing.T) {
	router, err := run()
	if err != nil {
		t.Fatal(err)
	}
	params := url.Values{}
	params.Add("key", "2")
	u := "/products/find?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		t.Log("http.NewRequest", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		logData(req, t)
	}
}

func logData(r *http.Request, t *testing.T) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	defer t.Fail()
	if err != nil {
		t.Logf("ioutil.ReadAll err %v", err)
		return
	}
	t.Logf("body: %v", string(body))
}
