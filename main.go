package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/koshigoe/monkeylang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
