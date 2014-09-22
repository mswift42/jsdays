package jsdays

import (
	"fmt"
	"html/template"
	"net/http"
)

// withLayout - take a template name and a templatefile
// and return it combined with layout.tmpl.
func withLayout(name, templ string) *template.Template {
	return template.Must(template.New(name).ParseFiles(templ, "templates/layout.tmpl"))
}

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
}

func home(w http.ResponseWriter, r *http.Request) {
	if err := withLayout("home", "templates/index.tmpl").Execute(w, map[string]string{"Pagetitle": "Task Overview"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if err := withLayout("about", "templates/about.tmpl").Execute(w, map[string]string{"Pagetitle": "About"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
