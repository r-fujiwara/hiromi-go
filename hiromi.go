package main

import (
  "./words"
  "fmt"
  "os"
)

// this is current stable on go run
func goHome() {
	fmt.Println("Go home")
}

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		if args[0] == "go" {
			words.SayHiromiGo()
		} else {
			goHome()
		}
	} else if len(args) == 0 {
		words.SayOrigin()
	} else {
		goHome()
	}
}
