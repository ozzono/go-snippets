package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	forms "github.com/synini/ifbra-forms"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/forms/1", form1)
	r.HandleFunc("/forms/2", form2)
	r.HandleFunc("/forms/3", form3)
	r.HandleFunc("/forms/4", form4)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the index page!")
}

func form1(w http.ResponseWriter, r *http.Request) {
	output, err := json.MarshalIndent(forms.Form1(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func form2(w http.ResponseWriter, r *http.Request) {
	output, err := json.MarshalIndent(forms.Form2(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func form3(w http.ResponseWriter, r *http.Request) {
	output, err := json.MarshalIndent(forms.Form3(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func form4(w http.ResponseWriter, r *http.Request) {
	output, err := json.MarshalIndent(forms.Form4(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
