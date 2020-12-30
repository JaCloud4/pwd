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

//Done. Serves homepage
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
  err:=r.ParseForm()
  if err!=nil{log.Fatal(err)}
  inputs:=r.Form //Parse form and extract slices in map
  exclude:=inputs["Exclusions"]//fmt.Println(reflect.TypeOf(inputs["Exclusions"])) //Slice of Integer, will test
  fmt.Println(len(exclude))//Will use lens(exclude)==4 {redirect}
  size, _:= strconv.Atoi(inputs["size"][0])  //fmt.Println(size)
  pw := app.RandPwdEx(size, exclude) //SUCCESS!  | ONly need to clear What if..all exclusions are marked
  pwd := password{
    Length: size,
    Code:   pw,
    Valid:  len(pw) >= 0}
  fmt.Println(pwd, "\n", pw)
  html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
  er := html.Execute(w, pwd)
  if er != nil {
    log.Fatal(er)
  }
}

func GoStart(_ string) {
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