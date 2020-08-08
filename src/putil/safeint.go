package sutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
