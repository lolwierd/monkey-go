package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/lolwierd/monkey-go/pkg/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! REPL commmminn!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
