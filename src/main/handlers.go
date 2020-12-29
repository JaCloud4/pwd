package main

import (
  "fmt"
  "github.com/JaCloud4/PwdMaster/pkg/app"
  "github.com/gorilla/mux"
  "html/template"
  "log"
  "math/rand"
  "net/http"
  "strconv"
  "time"
)

func pwdhome(w http.ResponseWriter, r *http.Request) {
  html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
  er := html.Execute(w, nil)
  if er != nil {
    log.Fatal(er)
  }
}
func pwdapi(w http.ResponseWriter, r *http.Request) {

}
func pwdresults(w http.ResponseWriter, r *http.Request) {
  type password struct {
    Length int
    Code   string
    Valid  bool
  }
  rand.Seed(time.Now().UTC().UnixNano())
  many := r.FormValue("size")
  //length, exclude := r.FormValue("size"), r.Form["Exclusions"]
  man, _ := strconv.Atoi(many)
  fmt.Println(man)
  pw := app.RandPwd(man)
  pwd := password{
    Length: man,
    Code:   pw,
    Valid:  len(pw) >= 0}
  fmt.Println(pwd, "\n", pw)
  html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
  er := html.Execute(w, pwd)
  if er != nil {
    log.Fatal(er)
  }
}

func setuproute() {
  r := mux.NewRouter()
  r.HandleFunc("/pwd", pwdhome).Methods("GET")
  r.HandleFunc("/pwd/api/{num}", pwdapi).Methods("GET")
  r.HandleFunc("/pwd/results", pwdresults).Methods("POST")

  r.HandleFunc("/pwd/", func(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request, "/pwd", 301)
  }).Methods("GET")
  r.HandleFunc("/pwd/{anythingelse}", func(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request, "/pwd", 301)
  }).Methods("GET") //Proofing url
  address := ":60952"
  fmt.Println("Program launched on localhost" + address)
  log.Fatal(http.ListenAndServe(address, r))
}