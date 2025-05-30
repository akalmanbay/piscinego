package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, v := range args {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
