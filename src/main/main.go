package main

import (
  "fmt"
  "github.com/QueseGen/PwdMaster/pkg/app"
  "math/rand"
  "time"
)

var numbers int
var listofInt []int

/*Running the Web Service, BASIC COMPLETED!
Currently working on Exclusion:
test.go> TestingExclusions()
pwdgen.go> RandPwdEx()
handlers.go> pwdresults()

 */
func main() {
  fmt.Println("http://localhost:60952/pwd")
  GoStart("n") //located in handlers.go
}
//Running a Console Version
func consoleplay()  {
  fmt.Println("Welcome to a PWD Generator")
  rand.Seed(time.Now().UTC().UnixNano()) // var ch1 chan string defer close(ch1) pwdd:= <-ch1
  for true {
    fmt.Println("How many numbers? Enter Zero to Quit...")
    reader := app.Intonly() // located in pkg/app
    if reader == 0 {
      break
    }
    pwdd := app.RandOrder(reader)
    fmt.Println(pwdd)
  }
  fmt.Println("Thank you for playing!!")
}


