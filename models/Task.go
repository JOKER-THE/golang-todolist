package models

type Task struct {
    Id          string
    Title       string
    Description string
}

func NewTask(id, title, description string) *Task {
    return &Task{id, title, description}
}