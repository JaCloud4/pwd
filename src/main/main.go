package main

import (
  "bufio"
  "fmt"
  "github.com/JaCloud4/PwdMaster/src/app"
  "github.com/gorilla/mux"
  "html/template"
  "log"
  "math/rand"
  "net/http"
  "os"
  "strings"
  "time"
)


var numbers int
var listofInt []int
func AndPwd(size int, p chan string) {
  for x:=0;x<size;x++{
    random := rand.Intn(5) + 1
    switch random {
    case 1:
     p<-app.RandInt()
    case 2:
      p<-app.RandCap()
    case 3:
      p<-app.RandLow()
    case 4:
      p<-app.RandSymbols()
    case 5:
      p<-app.RandChar()
    }}
}

//This version asks how many instances expected,
// then loops to capture integer entries only
func CreatesAfile(){
  fmt.Println("Name of File: ")
  reader := bufio.NewReader(os.Stdin)
  input, er := reader.ReadString('\n')
  if er !=nil {log.Fatal(er)}
  input= strings.TrimSpace(input)
 // printfile, fileerror:=sutil.Openfile(input)
 // fmt.Println(printfile,fileerror)
}
func RandPWDGenerator(){
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Random Number:", app.RandInt())
  fmt.Println("Random Uppered:",app.RandCap())
  fmt.Println("Random Lowered:",app.RandLow())
  fmt.Println("Random Symbol:",app.RandSymbols())
  fmt.Println("Random Mixed:",app.RandChar())
  //fmt.Println(RandPwd(12, nil))

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
func mai() {
  var ch1 chan string
  fmt.Println("Welcome to a PWD Generator")
  fmt.Println("How many numbers? Enter Zero to Quit...")
  reader :=app.Intonly()
  pwd:=app.RandPwd(reader,ch1)
  fmt.Println(pwd)
  setuproute()
}

func main()  {
 RandPWDGenerator()
}
