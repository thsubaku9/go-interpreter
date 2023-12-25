package main

import (
	"fmt"
	"monkey-i/repl"
	"os"
	"os/user"
)

func main() {
	userRef, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", userRef.Username)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
