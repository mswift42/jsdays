package jsdays

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
)

// Task consists of a Summary and the Task Content,
// both strings, the task status, "Done" or "TODO"
// and a scheduled Date.
type Task struct {
	Id        int64  `json:"id" datastore:"-"`
	Summary   string `json:"summary"`
	Content   string `json:"content" datastore:",noindex"`
	Status    string `json:"status"`
	Scheduled string `json:"scheduled"`
}

func defaultTaskList(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Task", "default_tasklist", 0, nil)
}

// withLayout - take a template name and a templatefile
// and return it combined with layout.tmpl.
func withLayout(name, templ string) *template.Template {
	return template.Must(template.New(name).ParseFiles(templ, "templates/layout.tmpl"))
}

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/newtask", newtask)
}

func home(w http.ResponseWriter, r *http.Request) {
	if err := withLayout("index", "templates/index.tmpl").Execute(w, map[string]string{"Pagetitle": "Task Overview"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if err := withLayout("about", "templates/about.tmpl").Execute(w, map[string]string{"Pagetitle": "About"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func newtask(w http.ResponseWriter, r *http.Request) {
	if err := withLayout("newtask", "templates/newtask.tmpl").Execute(w, map[string]string{"Pagetitle": "New Task"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
