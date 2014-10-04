package jsdays

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
	"strconv"
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

// keyForID - return the datastore.Key for a given IntID
func keyForID(c appengine.Context, id int64) *datastore.Key {
	return datastore.NewKey(c, "Task", "", id, defaultTaskList(c))
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
	http.HandleFunc("/edittask", edittask)
	http.HandleFunc("/updatetask", updatetask)
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
	t := Task{Summary: r.FormValue("formsummary"),
		Content:   r.FormValue("formcontent"),
		Scheduled: r.FormValue("formscheduled"),
		Status:    "TODO"}
	if _, err := t.save(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func edittask(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	id, _ := strconv.ParseInt(r.FormValue("taskid"), 10, 64)
	var edittask Task
	key := keyForID(c, id)
	if err := datastore.Get(c, key, &edittask); err != nil {
		panic(err)
	}
	withLayout("edittask", "templates/edittask.tmpl").Execute(w,
		map[string]interface{}{"Pagetitle": "Edit Tasks",
			"Summary": edittask.Summary,
			"Content": edittask.Content, "Taskid": id,
			"Scheduled": edittask.Scheduled,
			"Status":    edittask.Status})
}

// updatetask - retrieve task with id 'id'.
// If delete button is pressed, delete task in datastore
// else update it's contents.
func updatetask(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	id, err := strconv.ParseInt(r.FormValue("taskid"), 10, 64)
	status := r.FormValue("taskstatus")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var task Task
	key := keyForID(c, id)
	if err := datastore.Get(c, key, &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if btn := r.FormValue("submitbutton"); btn == "delete" {
		datastore.Delete(c, key)
	} else {
		task.Summary = r.FormValue("formsummary")
		task.Content = r.FormValue("formcontent")
		task.Status = r.FormValue("formstatus")
		task.Scheduled = r.FormValue("formscheduled")
		if status == "on" {
			task.Status = "DONE"
		}
		if _, err := task.save(c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
