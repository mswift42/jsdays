package jsdays

import (
	"appengine/aetest"
	"appengine/datastore"
	"strconv"
	"testing"
)

func TestEditTask(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	key := datastore.NewKey(c, "Task", "", 123456789, defaultTaskList(c))
	id, _ := strconv.ParseInt("123456789", 10, 64)
	if id != 123456789 {
		t.Fatal(err)
	}
	if _, err := datastore.Put(c, key, &Task{Summary: "some summary"}); err != nil {
		t.Fatal(err)
	}
	task1ret := Task{}
	newkey := keyForID(c, id)
	if err := datastore.Get(c, newkey, &task1ret); err != nil {
		t.Fatal(err)
	}
	tid := newkey.IntID()
	if task1ret.Summary != "some summary" {
		t.Error("Expected <some summary>, got: ", task1ret.Summary)
	}
	if tid != 123456789 {
		t.Error("Expected id 123456789, got: ", tid)
	}

}
