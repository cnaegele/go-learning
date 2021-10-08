package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	myEnv, exist := os.LookupEnv("GOPATH")
	if exist {
		fmt.Printf("Your Var %v", myEnv)

	}

}
