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

func TestKeyAndSaveAndList(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	t1 := Task{Summary: "some summary", Id: 123}
	t2 := Task{Summary: "second summary", Id: 222}
	if _, err := t1.save(c); err != nil {
		t.Fatal(err)
	}
	if _, err := t2.save(c); err != nil {
		t.Fatal(err)
	}
	k1 := t1.key(c)
	k1id := k1.IntID()
	if k1id != 123 {
		t.Error("Expected 123, got: ", k1id)
	}
	k2 := t2.key(c)
	k2id := k2.IntID()
	if k2id != 222 {
		t.Error("Expected 222, got: ", k2id)
	}
	tasks, err := listTasks(c)
	if err != nil {
		t.Fatal(err)
	}
	if tasks[0].Id != 123 {
		t.Error("Expected 123, got: ", tasks[0].Id)
	}
	if tasks[1].Id != 222 {
		t.Error("Expected 222, got: ", tasks[1].Id)
	}
}
