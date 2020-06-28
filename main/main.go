package main

import (
	"fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Helloworld")
}
func setuproute() {
	r:=mux.NewRouter()
	r.HandleFunc("/pwd", helloworld).Methods("GET")
	http.ListenAndServe(":60952", r)
}
func main() {
 // fmt.Println("Welcome to a PWD Generator")
	setuproute()
}
