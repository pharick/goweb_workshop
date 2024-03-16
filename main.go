package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pharick/cool_web_app/handlers"
)

func main() {
	// router is making choice based on the URL
	// which handler to call
	r := mux.NewRouter()

	// connect the URL to the handler
	r.HandleFunc("/", handlers.IndexPage)
	r.HandleFunc("/users/{username}", handlers.UserPage)

	// start the server
	http.ListenAndServe(":3000", r)
	// :3000 is the port number
}
