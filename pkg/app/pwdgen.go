package app

import (
  "fmt"
  "math/rand"
  "time"
)

type listofsym []int
var(
   Order []int = []int {}
   ExcludeL []int=[]int{1,2,3,4}
   listofsymbols []int = []int{64, 37, 43, 92, 47, 39, 33, 35, 36, 94, 63, 58, 44, 40, 41, 123, 125, 91, 93, 126, 45, 95, 46}//listofsymbols []rune=[]rune{'@', '%', '+','\\','/','\'','!','#','$','^','?',':',',','(',')','{','}','[',']','~','-','_','.'}
   )

//Password Generator with no Exclusions
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
//Password Generator with Exclusions
func RandPwdEx(size int, exclude []string) (string){
  var pwd string
  if len(exclude)==0 {
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
    }}} else if len(exclude)==4 {
      pwd=pwd+"Too Exclusive! Please use three or less criteria's."
  } else {
    pwd=TailoredPwd(size, exclude)
      }
  return pwd
}
func RandOrder(size int) (string){
  var pwd string
  order := []int{1,2,3,4,5}
  rand.Shuffle(len(order), func(i, j int) {
    order[i], order[j] =order[j], order[i]
  })
  for x := 0; x < size; x++ {
    rand.Seed(time.Now().UnixNano())
    if len(order)==0 {
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
func TailoredPwd(size int, cuts []string) (string){
  var pwd string
  var list []int
  if cuts[0]=="n" || cuts[0]=="u" || cuts[0]=="l" || cuts[0]=="s"{
    list=excludeApi(cuts)
   } else if cuts[0]=="Numbers" || cuts[0]=="Uppercase" || cuts[0]=="Lowercase" || cuts[0]=="Symbols" {
    list=excludeThese(cuts)
} else{}
  for x := 0; x < size; x++ {
    random := rand.Intn(len(list))
    switch list[random]{
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