package putil

import (
	"bufio"
	"log"
	"os"
)
var lines []string
func filestorage(file string)(error){
	return nil
}

func Openfile(file string) ([]string, error) {
openfil, err:=os.Open(file)
if err !=nil{log.Fatal(err)}
//	err=openfil.Chmod(755)
//	if err!=nil{ log.Fatal(err)}
	scanner := bufio.NewScanner(openfil)
	for scanner.Scan(){
		line:=scanner.Text()
		lines=append(lines, line)
	}
	err=openfil.Close()
	if err!=nil{log.Fatal(err)}
	if scanner.Err()!=nil{ return nil, scanner.Err()}
	return lines ,nil
}
