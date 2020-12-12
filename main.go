package main

import (
	"FoxLite/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Wellcome to FoxLite programming language!\n", user.Username)
	fmt.Printf("Feel free to type some Visual FoxPro commands like\n")
	repl.Start(os.Stdin, os.Stdout)
}
