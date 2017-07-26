package main

import (
	"fmt"

	"os"

	"github.com/scheshan/MarkdownHere/mh"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Getwd())

	e, _ := os.Executable()

	fmt.Println(e)

	server := mh.NewHTTPServer("", 5000)
	err := server.Start()

	if err != nil {
		fmt.Print(err)
	}
}
