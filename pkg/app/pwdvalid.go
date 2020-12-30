package app

import (
  "bufio"
  "fmt"
  "log"
  "math/rand"
  "os"
  "strconv"
  "strings"
  "time"
)
//Read Below
/*
Validators COMPLETE!
IntOnly> Can check and get Only Integer Values
Rand*> Get Random Integers,Uppercase,Lowercase, and Symbols
RandChar> Serves as a wildcard
 */
func Intonly() int {
  var intput int
  var noint error

  for true {
    reader := bufio.NewReader(os.Stdin)
    input, er := reader.ReadString('\n') // 5 3.3
    if er != nil {log.Fatal(er)}

    input = strings.TrimSpace(input)
    if input == "stop" {break}

    intput, noint = strconv.Atoi(input)
    if noint != nil {
      fmt.Println("Please enter a Integer: ")
    } else {break}
  }
  return intput
}

func RandInt() string {
  random := rand.Intn(57-48)+48 //ASCII 48-57 == '0-9'
  //	fmt.Println(random)
  return string(random)
}
func RandCap() string {
  random := rand.Intn(90-65) + 65 //ASCII 65-96 == 'A-Z'
  return string(random)
}
func RandLow() string {
  random := rand.Intn(122-97) + 97 //ASCII 97-122 == 'a-z'
  return string(random)
}
func RandSymbols() string {
  var list listofsym = listofsymbols
  random := rand.Intn(23)
  return string(list[random])
}
func RandChar() string {
  random := rand.Intn(122-65) + 65 //ASCII 65-122 == 'A-z'
  if random == 96 {
    RandChar()
  }
  return string(random)
} //A-z some symbols 91,92,93,94,95

//Testing Logic
func RandTestGenerator() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Random Number:", RandInt())
  fmt.Println("Random Uppered:", RandCap())
  fmt.Println("Random Lowered:", RandLow())
  fmt.Println("Random Symbol:", RandSymbols())
  fmt.Println("Random Mixed:", RandChar())
}

func excludeThese(a []string) ([]int){
  list:=map[string]int {"Numbers":1,"Uppercase":2,"Lowercase":3,"Symbols":4}
  var newlist []int
  for _, item := range a{
    delete(list, item)
  }
  for _,num:= range list{
    newlist=append(newlist,num)
  }
  return newlist
}
func excludeD(a []int, b []int) []int {
  // Turn b into a map
  var m map[int]bool
  m = make(map[int]bool, len(b))
  for _, s := range b {
    m[s] = false
  }
  // Append values from the longest slice that don't exist in the map
  var diff []int
  for _, s := range a {
    if _, ok := m[s]; !ok {
      diff = append(diff, s)
      continue
    }
    m[s] = true
  }
  // Sort the resulting slice
  return diff
}