package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep:= "", " "

	for idx, arg := range(os.Args[1:]) {
		s = strconv.Itoa(idx) + sep + arg		
		fmt.Println(s)
	}
}

