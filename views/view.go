package views

import "html/template"

// NewView returns a view
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View is a view type
type View struct {
	Template *template.Template
}
