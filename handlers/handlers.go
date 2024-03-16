package handlers

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

var users = map[string]map[string]any{
	"artem": {
		"Name": "Artem",
		"City": "Bangkok",
	},

	"john": {
		"Name": "John",
		"City": "New York",
	},

	"jane": {
		"Name": "Jane",
		"City": "Los Angeles",
	},
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, ok := users[username]
	log.Printf("User: %v, ok: %v", user, ok)

	if !ok {
		http.NotFound(w, r)
		return
	}

	renderTemplate(w, "user.html", user)
}
