package main

import (
	"log"
	myServer "net-http/cmd/server"
)

func main() {
	err := myServer.StartVlanServer()
	if err != nil {
		log.Fatal(err)
	}

}
