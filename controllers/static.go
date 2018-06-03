package controllers

import (
	"github.com/clarakko/lenslocked/views"
)

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bulma", "views/static/home.gohtml"),
		Contact: views.NewView("bulma", "views/static/contact.gohtml"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
}
