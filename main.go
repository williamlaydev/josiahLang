package main

import (
	"fmt"
	"josiahLang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the REPL for JosiahLang\n",
		user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
