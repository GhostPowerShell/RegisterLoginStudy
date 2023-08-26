package main

import (
	"net/http"

	"github.com/GhostPowerShell/RegisterLoginStudy/handlers"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", handlers.LoginPageHandler) // GET

	router.HandleFunc("/index", handlers.IndexPageHandler) // GET
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	router.HandleFunc("/register", handlers.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
