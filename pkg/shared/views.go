package shared

import (
	"html/template"
)

func ParseTemplate(name string) *template.Template {
	return template.Must(template.ParseFiles("templates/base.html", "templates/" + name + ".html"))
}