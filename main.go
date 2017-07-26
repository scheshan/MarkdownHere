package main

import (
	"github.com/scheshan/SimpleServer/server"
)

func main() {
	var ser = server.HTTPServer{}
	ser.Start()
}
