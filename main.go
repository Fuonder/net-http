package main

import (
	"log"
	myServer "net-http/cmd/server"
)

func main() {
	err := myServer.StartAuthPageServer()
	if err != nil {
		log.Fatal(err)
	}

}
