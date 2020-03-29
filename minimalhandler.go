package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	form "github.com/synini/ifbra-forms"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/form1.json", form1).Methods("GET")
	r.HandleFunc("/form2.json", form2).Methods("GET")
	r.HandleFunc("/form3.json", form3).Methods("GET")
	r.HandleFunc("/form4.json", form4).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling index")
	fmt.Fprint(w, "This is the index page!")
}

func form1(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling form1")
	output, err := json.MarshalIndent(form.Form1(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func form2(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling form2")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	output, err := json.MarshalIndent(form.Form2(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func form3(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling form3")
	output, err := json.MarshalIndent(form.Form3(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func form4(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling form4")
	output, err := json.MarshalIndent(form.Form4(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func CORS(h http.HandlerFunc) http.Handler {
	allowHeaders := strings.Join([]string{
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Origin",
		"Authorization",
		"Content-Type",
		"Cookie",
		"Set-Cookie",
		"X-Login-With",
		"X-Requested-With",
	}, ",")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Headers", allowHeaders)
		w.Header().Add("Access-Control-Expose-Headers", allowHeaders)
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
