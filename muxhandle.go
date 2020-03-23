package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	form "github.com/synini/ifbra-forms"
)

func main() {
	mux := http.NewServeMux()

	index := http.HandlerFunc(indexHandler)
	form1 := http.HandlerFunc(form1)
	form2 := http.HandlerFunc(form2)
	form3 := http.HandlerFunc(form3)
	form4 := http.HandlerFunc(form4)

	mux.Handle("/", index)
	mux.Handle("/forms/1", form1)
	mux.Handle("/forms/2", form2)
	mux.Handle("/forms/3", form3)
	mux.Handle("/forms/4", form4)

	log.Println("Listening...")
	http.ListenAndServe(":8000", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling index page")
	fmt.Fprint(w, "This is the index page!")
}

func form1(w http.ResponseWriter, r *http.Request) {
	log.Println("Returning data from form1")
	output, err := json.MarshalIndent(form.Form1(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func form2(w http.ResponseWriter, r *http.Request) {
	log.Println("Returning data from form2")
	output, err := json.MarshalIndent(form.Form2(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func form3(w http.ResponseWriter, r *http.Request) {
	log.Println("Returning data from form3")
	output, err := json.MarshalIndent(form.Form3(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func form4(w http.ResponseWriter, r *http.Request) {
	log.Println("Returning data from form4")
	output, err := json.MarshalIndent(form.Form4(), "", "	")
	if err != nil {
		log.Printf("marshal err: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
