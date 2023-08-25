package main

import (
	"fmt"
	"net/http"

	h "github.com/GhostPowerShell/helpers"
)

func main() {

	username, email, passwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username = r.FormValue("username")
		email = r.FormValue("email")
		passwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		usernameCheck := h.IsEmpty(username)
		emailCheck := h.IsEmpty(email)
		passwdCheck := h.IsEmpty(passwd)
		pwdConfirmCheck := h.IsEmpty(pwdConfirm)

		if usernameCheck || emailCheck || passwdCheck || passwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
		}
		if passwd == pwdConfirm {
			fmt.Fprintln(w, "Registration successful.")
		} else {
			fmt.Fprintln(w, "Password must be the same.")
		}
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		passwd = r.FormValue("password")

		emailCheck := h.IsEmpty(email)
		passwdCheck := h.IsEmpty(passwd)

		if emailCheck || passwdCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
		}

		dbPwd := "mouse"   // DB simulation
		dbEmail := "mouse" // DB simulation

		if email == dbEmail && passwd == dbPwd {
			fmt.Fprintln(w, "Login succesful!")
		} else {
			fmt.Fprintln(w, "Login failed!")
		}
		http.ListenAndServe(":8080", mux)
	})
}
