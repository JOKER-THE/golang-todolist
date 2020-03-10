package models

/**
 * Task - модель, представляющая задачу в ToDoList
 *
 * Id - уникальный идентификатор, генерируемый
 * отдельной функцией GenerateId()
 * Title - заголовок задачи в ToDoList
 * Description - описание задачи в ToDoList
 */
type Task struct {
    Id          string
    Title       string
    Description string
}

/**
 * NewTask - конструктор структуры Task,
 * возвращающий объект задачи
 */
func NewTask(id, title, description string) *Task {
    return &Task{id, title, description}
}