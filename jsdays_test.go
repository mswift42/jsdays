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
	key := datastore.NewKey(c, "Task", "", 1, nil)
	id, _ := strconv.ParseInt("123456789", 10, 64)
	if id != 123456789 {
		t.Fatal(err)
	}
	if _, err := datastore.Put(c, key, &Task{Summary: "some summary", Id: 123456789}); err != nil {
		t.Fatal(err)
	}
	task1ret := Task{}
	if err := datastore.Get(c, key, &task1ret); err != nil {
		t.Fatal(err)
	}
	if task1ret.Summary != "some summary" {
		t.Error("Expected <some summary>, got: ", task1ret.Summary)
	}
	if task1ret.Id != 123456789 {
		t.Error("Expected <123456789>, got: ", task1ret.Id)
	}

}
