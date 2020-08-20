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

type listofsym []int
var(
   Order []int = []int {1,2,3,4,5}
   Alpha=map[string]int {"Numbers":1,"Uppercase":2,"Lowercase":3,"Symbols":4,"Extra":5}
   Beta=map[int]string {1:"Numbers",2:"Uppercase",3:"Lowercase",4:"Symbols",5:"Extra"}
   listofsymbols []int = []int{64, 37, 43, 92, 47, 39, 33, 35, 36, 94, 63, 58, 44, 40, 41, 123, 125, 91, 93, 126, 45, 95, 46}//listofsymbols []rune=[]rune{'@', '%', '+','\\','/','\'','!','#','$','^','?',':',',','(',')','{','}','[',']','~','-','_','.'}
   )

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

func RandPwd(size int) (string){
  var pwd string
	for x := 0; x < size; x++ {
		random := rand.Intn(5) + 1
		switch random {
		case 1:
			pwd+= RandInt()
		case 2:
			pwd+= RandCap()
		case 3:
			pwd+= RandLow()
		case 4:
			pwd+= RandSymbols()
		case 5:
			pwd+= RandChar()
		default:
			fmt.Errorf("Where Waldo?")
		}
	}
	return pwd
}
func RandOrder(size int) (string){
  var pwd string
  order := Order
  rand.Shuffle(len(order), func(i, j int) {
    order[i], order[j] =order[j], order[i]
  })
  for x := 0; x < size; x++ {
    rand.Seed(time.Now().UnixNano())
    if len(order)==0 {
      order=Order
      rand.Shuffle(len(order), func(i, j int) {
        order[i], order[j] = order[j], order[i]
      })
    }
    fmt.Println(order)
    switch order[0] {
    case 1:
      pwd+= RandInt()
      order=order[1:]
    case 2:
      pwd+= RandCap()
      order=order[1:]
    case 3:
      pwd+= RandLow()
      order=order[1:]
    case 4:
      pwd+= RandSymbols()
      order=order[1:]
    case 5:
      pwd+= RandChar()
      order=order[1:]
    default:
      fmt.Errorf("Where Waldo?")
    }
  }
  return pwd
}
func RandomAlpha(size int, cuts []string) (string){
  var pwd string

  alpha:=Alpha
  for x:=0; x<len(cuts);x++{
    delete(alpha, cuts[x])
  }

  var order []int
  for _, orders := range alpha {
    order = append(order, orders)
  }

  for x := 0; x < size; x++ {
    rand.Seed(time.Now().UnixNano())
    if len(order)==0 {
      order=Order
      rand.Shuffle(len(order), func(i, j int) {
        order[i], order[j] = order[j], order[i]
      })
    }
    fmt.Println(order)
    switch order[0] {
    case 1:
      pwd+= RandInt()
      order=order[1:]
    case 2:
      pwd+= RandCap()
      order=order[1:]
    case 3:
      pwd+= RandLow()
      order=order[1:]
    case 4:
      pwd+= RandSymbols()
      order=order[1:]
    case 5:
      pwd+= RandChar()
      order=order[1:]
    default:
      fmt.Errorf("Where Waldo?")
    }
  }
  return pwd
}