package server

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func StartVMServer() error {
	http.Handle("/", Conveyor(http.HandlerFunc(rootHandle)))
	err := http.ListenAndServe("192.168.0.157:8080", nil)
	if err != nil {
		return fmt.Errorf("VM server start: %w", err)
	}
	return nil
}
