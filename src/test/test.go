package test

import (
  "fmt"
  "github.com/QueseGen/PwdMaster/pkg/app"
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
    //pwdd := app.Rand(reader, []string{"Uppercase", "Lowercase", "Symbols"})
   // fmt.Println(pwdd)
  }
  fmt.Println("Thank you for playing!!")
}
func APIPrint(t *testing.T)  {
  rand.Seed(time.Now().UTC().UnixNano()) // var ch1 chan string defer close(ch1) pwdd:= <-ch1
  fmt.Println("Welcome to a PWD Generator")
  var tests = []struct{
    input string
    expected int
  }{}
    for _,test := range tests{
      fmt.Println(test)
    }
    fmt.Println(tests)
  for true {
    fmt.Println("How many numbers? Enter Zero to Quit...")
}
}
func NetwongerTest(t *testing.T){
  //make call to API, enter values

  /* pwdvalid>Newtonger
  func Intonly(input string) int {
    var intput int
    var noint error
  if input=="" {
    for true {
      reader := bufio.NewReader(os.Stdin)
      input, er := reader.ReadString('\n') // 5 3.3
      if er != nil {log.Fatal(er)}

      input = strings.TrimSpace(input)
      if input == "stop" {break}

      intput, noint = strconv.Atoi(input)
      if noint != nil {
        fmt.Println("Please enter a Integer: ")
      }}} else if input == "1" {

      }
      return intput
  }
   */
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
