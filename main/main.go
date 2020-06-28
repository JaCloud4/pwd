package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Helloworld")
}
func setuproute() {
	//_:=mux.NewRouter()
	http.HandleFunc("/api/home", helloworld)
	http.ListenAndServe(":63342", nil)
}
func main() {
	setuproute()
}
