package main

import (
  "fmt"
  "github.com/JaCloud4/PwdMaster/pkg/app"
  "github.com/gorilla/mux"
  "html/template"
  "log"
  "math/rand"
  "net/http"
  "testing"
  "time"
)

//var numbers int
//var listofInt []int

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

func pwdhome(w http.ResponseWriter, r *http.Request) {
  //fp := path.Join("templates", "index.html")
	html, _ := template.ParseFiles("/Users/quese/Documents/PwdMaster/src/templates/pwdhome.html")
	er:=html.Execute(w, nil)
	if er!=nil{log.Fatal(er)}
}
func setuproute() {
	r := mux.NewRouter()
	r.HandleFunc("/pwd", pwdhome).Methods("GET")
	log.Fatal(http.ListenAndServe(":60952", r))
}
func main() {
  rand.Seed(time.Now().UTC().UnixNano())
 // var ch1 chan string defer close(ch1) pwdd:= <-ch1
  fmt.Println("Welcome to a PWD Generator")  //setuproute()
  for true{
  fmt.Println("How many numbers? Enter Zero to Quit...")
  reader := app.Intonly()
  if reader==0{break}
  pwdd:= app.RandPwd(reader)
  fmt.Println(pwdd)
  }
  fmt.Println("Thank you for playing!!")
}

func successfullyrandomapis()  {
 RandPWDGenerator()
}
func RandPWDGenerator(){
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Random Number:", app.RandInt())
  fmt.Println("Random Uppered:", app.RandCap())
  fmt.Println("Random Lowered:", app.RandLow())
  fmt.Println("Random Symbol:", app.RandSymbols())
  fmt.Println("Random Mixed:", app.RandChar())
}