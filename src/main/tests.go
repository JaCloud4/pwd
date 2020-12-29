package main

import (
  "fmt"
  "github.com/JaCloud4/PwdMaster/pkg/app"
  "math/rand"
  "testing"
  "time"
)

func BasicCompTest() {
  app.RandTestGenerator()
}
/*Current Tasks:
1. Set conditions in console(pwdgen.go> RandPwdEx())
2. Implement via HTTP(handlers.go> pwdresults())
 */
func TestingExclusions() {
  rand.Seed(time.Now().UTC().UnixNano()) // var ch1 chan string defer close(ch1) pwdd:= <-ch1
  fmt.Println("Welcome to a PWD Generator")
  for true {
    fmt.Println("How many numbers? Enter Zero to Quit...")
    reader := app.Intonly()
    if reader == 0 {
      break
    }
    //pwdd:= app.RandOrder(reader)
    pwdd := app.RandomAlpha(reader, []string{"Numbers", "Symbols"})
    fmt.Println(pwdd)
  }
  fmt.Println("Thank you for playing!!")
}

func POAusingChannels(t *testing.T) {
  var size int
  var p chan string
  for x := 0; x < size; x++ {
    random := rand.Intn(5) + 1
    switch random {
    case 1:
      p <- app.RandInt()
    case 2:
      p <- app.RandCap()
    case 3:
      p <- app.RandLow()
    case 4:
      p <- app.RandSymbols()
    case 5:
      p <- app.RandChar()
    default:
      t.Errorf("Not so reandom", random)
    }
  }
  close(p)
  _, er := <-p
  if er != false {
    t.Errorf("Something went wrong with channel", p)
  }
}
