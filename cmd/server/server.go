package main

import "net/http"

func main() {
	err := http.ListenAndServe(`26.53.94.227:8080`, nil)
	if err != nil {
		panic(err)
	}
}
