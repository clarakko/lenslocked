package main

import (
	"net/http"

	"github.com/clarakko/lenslocked/views"
	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func main() {
	homeView = views.NewView("bulma", "views/home.gohtml")
	contactView = views.NewView("bulma", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
