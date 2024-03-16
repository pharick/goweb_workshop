package main

import (
	"fmt"
	"net/http"

	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/pharick/cool_web_app/database"
	"github.com/pharick/cool_web_app/handlers"
)

func main() {
	// connect to the database
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	db, err := database.Connect(connString)
	if err != nil {
		// print error and exit
		log.Fatal(err)
	}

	// create a new App object
	app := handlers.NewApp(db)

	// router is making choice based on the URL
	// which handler to call
	r := mux.NewRouter()

	// connect the URL to the handler
	r.HandleFunc("/", app.IndexPage)
	r.HandleFunc("/users/{username}", app.UserPage)

	log.Printf("Server starting on port 3000")
	// start the server
	http.ListenAndServe(":3000", r)
	// :3000 is the port number
}
