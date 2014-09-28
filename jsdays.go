package jsdays

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

// Task consists of a Summary and the Task Content,
// both strings, the task status, "Done" or "TODO"
// and a scheduled Date.
type Task struct {
	Summary   string
	Content   string
	Status    string
	Scheduled string
}

func createTable(t Task) {
	db, err := gorm.Open("sqlite3", "days.db")
	if err != nil {
		panic(err)
	}
	db.CreateTable(&t)
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
