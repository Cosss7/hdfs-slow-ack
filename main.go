package main

import (
	"example/user/hello/cmd"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	var err error
	if os.Args[1] == "s" {
		err = cmd.Server()
	} else if os.Args[1] == "c" {
		host := os.Args[2]
		err = cmd.Client(host)
	}

	if err != nil {
		fmt.Println(err)
	}
}
