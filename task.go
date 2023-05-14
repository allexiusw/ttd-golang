package store

import "strconv"

// Task job to be done or completed
type Task struct {
	ID    int    // identifier of the task
	Title string // Title of the task
	Done  bool   // If task is completed or not
}

// Datastore manages a list of tasks stored in memory
type Datastore struct {
	tasks []Task
}

// GetPendingTasks returns all the tasks which need to be done
func (ds *Datastore) GetTasks(status string) []Task {
	if status == "nil" {
		return ds.tasks
	}
	var bool_status, _ = strconv.ParseBool(status)
	var pendingTasks []Task
	for _, task := range ds.tasks {
		if task.Done == bool_status {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}
