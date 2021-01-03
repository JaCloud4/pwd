package main

import (
  "encoding/json"
  "fmt"
  "github.com/JaCloud4/PwdMaster/pkg/app"
  "github.com/gorilla/mux"
  "html/template"
  "log"
  "math/rand"
  "net/http"
  "strconv"
  "strings"
  "time"
)

type password struct {
  Length int `json:"length"`
  Message   string `json:"credentials"`
  Exclusions []string `json:"criteria"`

}

//Done. Serves homepage
func pwdhome(w http.ResponseWriter, r *http.Request) {
  html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
  er := html.Execute(w, nil)
  if er != nil {
    log.Fatal(er)
  }
}
func pwdresults(w http.ResponseWriter, r *http.Request) {
  rand.Seed(time.Now().UTC().UnixNano())
  err:=r.ParseForm()
  if err!=nil{log.Fatal(err)}
  inputs:=r.Form //Parse form and extract slices in map
  exclude:=inputs["Exclusions"]//fmt.Println(reflect.TypeOf(inputs["Exclusions"])) //Slice of Integer, will test
  size, _:= strconv.Atoi(inputs["size"][0])  //fmt.Println(size)
  pw := app.RandPwdEx(size, exclude) //PASSWORD CREATED
  pwd := password{ //Persistency and JSON capability
    Length: size ,
    Message:   pw,
    Exclusions:  exclude}
  fmt.Println(pwd, "\n", pw)
  html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
  er := html.Execute(w, pwd)
  if er != nil {
    log.Fatal(er)
  }
}
func pwdapi(w http.ResponseWriter, r *http.Request) {
  //PWDJSON='[{}]'
  vars := mux.Vars(r)
  size := app.IntApi(vars["size"])
  exclude := strings.Split(vars["Exclusions"], "") //need int
  pw := app.RandPwdEx(size, exclude) //PASSWORD CREATED

  fmt.Println("Size: ", size)
  fmt.Println("Exclusions: ", exclude)
  switch r.Method {
  case http.MethodGet:
    fmt.Println(vars, "\n", r.Method)
    if size != 0 {
      pwd := password{ //Persistency and JSON capability
        Length:     size,
        Message:    pw,
        Exclusions: exclude}
      fmt.Println("Password! ", pwd.Message, "\n-----------STRUC-----------\n", pwd)
      marshal, _ := json.Marshal(&pwd)
      err:=json.Unmarshal(marshal, &pwd)
      if err != nil{
        log.Fatal(err)
      }
      fmt.Println("First Marshalled: \n", string(marshal), "\n\n Thenn UnMarshalled:\n", &pwd )
      w.Header().Set("Content-Type", "application/json")
      w.Write(marshal)
    } else {
    html, _ := template.ParseFiles("/Users/quese/go/src/github.com/JaCloud4/PwdMaster/src/templates/tryindex.html")
      pwd := password{ //Persistency and JSON capability
        Length:     size,
        Message:    pw,
        Exclusions: exclude}
    er := html.Execute(w, pwd)
    if er != nil {
      log.Fatal(er)
    }
  }
    //exclude:=vars["ulsn"]
    /* pwd := password{ //Persistency and JSON capability
      Length: size ,
      Message:   pw,
      Exclusions:  exclude}
    fmt.Println("Type: ",reflect.TypeOf(size),
                  "\nLength: ", size,
                  "\nExclusions: ", exclude)
    */
    /* pw := app.RandPwdEx(size, exclude) //PASSWORD CREATED
    pwd := password{ //Persistency and JSON capability
      Length: size ,
      Message:   pw,
      Exclusions:  exclude} */
  }
}

func GoStart(_ string) {
  r := mux.NewRouter()
  r.HandleFunc("/pwd", pwdhome).Methods("GET")
  r.HandleFunc("/pwd/", func(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request, "/pwd", 301)}).Methods("GET")

  r.HandleFunc("/pwd/results", pwdresults).Methods("POST")

  r.HandleFunc("/pwd/api/{size}/{ulsn}", pwdapi).Methods("GET")

  //NEED REDIRECTION
  r.HandleFunc("/pwd/{anything}", func(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request, "http://localhost:60952/pwd", 301)}).Methods("GET") //Proofing url
    log.Fatal(http.ListenAndServe(":60952", r))


}