package main

import (
	"log"
	"net/http"

	"github.com/PaddleHQ/go-interview/internal/products"
	"github.com/gorilla/mux"
)

func main() {
	router, err := run()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Starting the server on port 8080")
	if err := http.ListenAndServe(":8080", Middleware(router)); err != nil {
		log.Fatal(err)
	}
}

func run() (*mux.Router, error) {
	productHandler := products.Handler{}

	r := mux.NewRouter()
	r.HandleFunc("/products/create", productHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/products/list", productHandler.List).Methods(http.MethodGet)
	r.HandleFunc("/products/find", productHandler.Find).Methods(http.MethodGet)
	return r, nil
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		h.ServeHTTP(w, r)
	})
}
