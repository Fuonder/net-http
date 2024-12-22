package main

import (
	"log"
	myServer "net-http/cmd/server"
)

func main() {
	err := myServer.StartVMServer()
	if err != nil {
		log.Fatal(err)
	}

}
