package controllers

import (
	"github.com/clarakko/lenslocked/views"
)

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bulma", "static/home"),
		Contact: views.NewView("bulma", "static/contact"),
		Faq:     views.NewView("bulma", "static/faq"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}
