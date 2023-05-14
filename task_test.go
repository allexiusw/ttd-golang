package store

import (
	"reflect"
	"testing"
)

func TestGetPendingTasks(t *testing.T) {
	t.Log("GetPendingTasks")

	ds := Datastore{
		tasks: []Task{
			{1, "Do housework", true},
			{2, "Buy milk", false},
		},
	}

	want := []Task{ds.tasks[1]}

	t.Log("should return the tasks which need to be completed")

	if got := ds.GetTasks(); !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v wanted %v", got, want)
	}
}

func TestGetDoneTasks(t *testing.T) {
	t.Log("GetDoneTasks")

	ds := Datastore{
		tasks: []Task{
			{1, "Do housework", true},
			{2, "Buy milk", false},
		},
	}

	want := []Task{ds.tasks[0]}

	t.Log("should return the tasks which need to be completed")

	if got := ds.getDoneTasks(); !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v wanted %v", got, want)
	}
}

func TestAllTasks(t *testing.T) {
	t.Log("GetAllTasks")

	ds := Datastore{
		tasks: []Task{
			{1, "Do housework", true},
			{2, "Buy milk", false},
		},
	}

	want := ds.tasks

	t.Log("should return the tasks which need to be completed")

	if got := ds.getAllTasks(); !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v wanted %v", got, want)
	}
}
