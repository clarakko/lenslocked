package controllers

import (
	"fmt"
	"net/http"

	"github.com/clarakko/lenslocked/views"
)

// NewUsers stuff
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bulma", "users/new")}
}

// Users is the users controller
type Users struct {
	NewView *views.View
}

// New is used to render the form where a user can create a new user account.
// Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when user creates new user account
// Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
