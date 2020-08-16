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
  "testing"
  "time"
)

func pwdhome(w http.ResponseWriter, r *http.Request) {
  html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
  er:=html.Execute(w, nil)
  if er!=nil{log.Fatal(er)}
}
func pwdapi(w http.ResponseWriter, r *http.Request) {
//mard
}


func pwdresults(w http.ResponseWriter, r *http.Request) {
  type password struct {
    Length int
    Code string
    Valid bool}

  rand.Seed(time.Now().UTC().UnixNano())
  many:=r.FormValue("size")
  man, _ := strconv.Atoi(many)
  fmt.Println(man)
  pw:=app.RandPwd(man)
  pwd:=password{
    Length: man,
    Code: pw,
    Valid: len(pw)>=0,}
  fmt.Println(pwd,"\n",pw)
  html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
  er:=html.Execute(w, pwd)
  if er!=nil{log.Fatal(er)}
}
func setuproute() {
	r := mux.NewRouter()
	r.HandleFunc("/pwd", pwdhome).Methods("GET")
  r.HandleFunc("/pwd/api/{num}", pwdapi).Methods("GET")
  r.HandleFunc("/pwd/results", pwdresults).Methods("POST")

  r.HandleFunc("/pwd/", func(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request,"/pwd", 301)
  }).Methods("GET")
  r.HandleFunc("/pwd/{anythingelse}", func(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request,"/pwd", 301)
  }).Methods("GET") //Proofing url
  address:=":60952"
  fmt.Println("Program launched on localhost"+address)
	log.Fatal(http.ListenAndServe(address, r))
}

func main() {
  setuproute()
  rand.Seed(time.Now().UTC().UnixNano())  // var ch1 chan string defer close(ch1) pwdd:= <-ch1
  fmt.Println("Welcome to a PWD Generator")
  for true{
  fmt.Println("How many numbers? Enter Zero to Quit...")
  reader := app.Intonly()
  if reader==0{break}
  pwdd:= app.RandPwd(reader)
  fmt.Println(pwdd)
  }
  fmt.Println("Thank you for playing!!")
}

func RandPWDGenerator(){
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Random Number:", app.RandInt())
  fmt.Println("Random Uppered:", app.RandCap())
  fmt.Println("Random Lowered:", app.RandLow())
  fmt.Println("Random Symbol:", app.RandSymbols())
  fmt.Println("Random Mixed:", app.RandChar())
}
func successfullyrandomapis()  {
  RandPWDGenerator()
}
func TestAndPwd(t *testing.T) {
  var size int
  var p chan string

  for x:=0;x<size;x++{
    random := rand.Intn(5) + 1
    switch random {
    case 1:
      p<- app.RandInt()
    case 2:
      p<- app.RandCap()
    case 3:
      p<- app.RandLow()
    case 4:
      p<- app.RandSymbols()
    case 5:
      p<- app.RandChar()
    default:
      t.Errorf("Not so reandom", random)
    }}
  close(p)
  _, er:=<-p
  if er!=false{t.Errorf("Something went wrong with channel", p)}
}
