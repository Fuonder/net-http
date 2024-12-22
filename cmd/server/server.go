package server

import (
	"fmt"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://yandex.ru/", http.StatusMovedPermanently)
}

func StartVMServer() error {
	http.Handle("/dummy", http.RedirectHandler("https://google.com", http.StatusMovedPermanently))
	http.HandleFunc("/search/", redirect)
	err := http.ListenAndServe("192.168.0.157:8080", nil)
	if err != nil {
		return fmt.Errorf("VM server start: %w", err)
	}
	return nil
}
