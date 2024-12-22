package server

import (
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ================\r\n"
	err := req.ParseForm()
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("API PAGE!"))
}

type myHandler struct{}

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello!")
	res.Write(data)
}

func StartAllEthServer() error {
	var h myHandler

	err := http.ListenAndServe(":8080", h)
	if err != nil {
		return fmt.Errorf("localhost Server start: %w", err)
	}
	return nil

}

func StartVlanServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc(`/api/`, apiPage)
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe("192.168.0.157:8080", mux)
	if err != nil {
		return fmt.Errorf("VM server start: %w", err)
	}
	return nil
}
