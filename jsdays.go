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

// defaultTaskList - return a new datastore key for a given
// http request context.
func defaultTaskList(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Task", "default_tasklist", 0, nil)
}

func (t *Task) key(c appengine.Context) *datastore.Key {
	if t.Id == 0 {
		return datastore.NewIncompleteKey(c, "Task", defaultTaskList(c))
	}
	return datastore.NewKey(c, "Task", "", t.Id, defaultTaskList(c))
}

func (t *Task) save(c appengine.Context) (*Task, error) {
	k, err := datastore.Put(c, t.key(c), t)
	if err != nil {
		return nil, err
	}
	t.Id = k.IntID()
	return t, nil
}
func listTasks(c appengine.Context) ([]Task, error) {
	tasks := []Task{}
	ks, err := datastore.NewQuery("Task").Ancestor(defaultTaskList(c)).Order("Status").Order("-Scheduled").GetAll(c, &tasks)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tasks); i++ {
		tasks[i].Id = ks[i].IntID()
	}
	return tasks, nil
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
	http.HandleFunc("/savetask", savetask)
}

func home(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	tasks, err := listTasks(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := withLayout("index", "templates/index.tmpl").Execute(w, map[string]interface{}{"Pagetitle": "Task Overview", "tasks": tasks}); err != nil {
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
func savetask(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(c, "Task", defaultTaskList(c))
	t := Task{Summary: r.FormValue("formsummary"),
		Content:   r.FormValue("formcontent"),
		Scheduled: r.FormValue("formscheduled"),
		Status:    "TODO"}
	if _, err := datastore.Put(c, key, &t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
