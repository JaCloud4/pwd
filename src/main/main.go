package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func pwdhome(w http.ResponseWriter, r *http.Request) {
	html, _ := template.ParseFiles("/Users/quese/Documents/PwdMaster/src/main/templates/pwdhome.html")
	log.Fatal(html.Execute(w, nil))
}
func setuproute() {
	r := mux.NewRouter()
	r.HandleFunc("/pwd", pwdhome).Methods("GET")
	log.Fatal(http.ListenAndServe(":60952", r))
}
func main() {
	fmt.Println("Welcome to a PWD Generator")
	setuproute()
}
