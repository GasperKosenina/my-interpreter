package main

import (
	"fmt"
	"os"
	"os/user"

	"interpreter-go.com/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, feel free to write in commands\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
