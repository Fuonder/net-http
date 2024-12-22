package server

import (
	"fmt"
	"io"
	"net/http"
)

// This is login test form, which handles method POST.
const form = `<html>
	<head>
	<title></title>
	</head>
	<body>
		<form action="/" method="post">
			<label>login <input type="text" name="login"></label>
			<label>password <input type="password" name="password"></label>
			<input type="submit" value="Login">
		</form>
	</body>
</html>`

func Auth(login, password string) bool {
	return login == `guest` && password == `demo`
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Welcome!")
		} else {
			http.Error(w, "Wrong login or password", http.StatusUnauthorized)
		}
		return
	} else {
		io.WriteString(w, form)
	}
}

func StartAuthPageServer() error {
	//mux := http.NewServeMux()
	//mux.HandleFunc(`/`, loginPage)
	err := http.ListenAndServe(`192.168.0.157:8080`, http.HandlerFunc(loginPage))
	if err != nil {
		return fmt.Errorf("auth sevrer %w", err)
	}
	return nil
}
