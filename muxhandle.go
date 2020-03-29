package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	form "github.com/synini/ifbra-forms"
	"google.golang.org/appengine"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling index page")
	fmt.Fprint(w, "This is the index page!")
	fmt.Fprint(w, "Browser thru the forms using the following paths:\n")
	fmt.Fprint(w, "https://ifbra-forms.appspot.com/forms/1\n")
	fmt.Fprint(w, "https://ifbra-forms.appspot.com/forms/2\n")
	fmt.Fprint(w, "https://ifbra-forms.appspot.com/forms/3\n")
	fmt.Fprint(w, "https://ifbra-forms.appspot.com/forms/4\n")
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

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/form1.json", form1).Methods("GET")
	r.HandleFunc("/form2.json", form2).Methods("GET")
	r.HandleFunc("/form3.json", form3).Methods("GET")
	r.HandleFunc("/form4.json", form4).Methods("GET")

	http.Handle("/", r)
}

func main() {
	appengine.Main()
}
