package models

type Task struct {
	ID   int
	Name string
	Done bool
}

func (myTask *Task) Toggle() {
	myTask.Done = !myTask.Done
}
