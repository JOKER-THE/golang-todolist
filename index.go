package main

import (
    "./models"
    "crypto/rand"
    "fmt"
    "html/template"
    "net/http"
)

var tasks map[string]*models.Task

/**
 * Метод генерации уникального идентификатора
 * задачи
 *
 * Возвращает случайно сгенерируемую 16-байтную строку
 */
func GenerateId() string {
    b := make([]byte, 16)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}

/**
 * Главная страница ToDoList
 *
 */
func indexHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("views/index.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "index", tasks)
}

/**
 * Страница создания новой задачи в ToDoList
 *
 */
func createHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("views/form.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "form", nil)
}

/**
 * Страница редактирования задачи в ToDoList
 *
 */
func updateHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("views/form.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    id := r.FormValue("id")
    task, found := tasks[id]
    if !found {
        http.NotFound(w, r)
    }
    t.ExecuteTemplate(w, "form", task)
}

/**
 * Метод сохранения данных задачи
 *
 * Считываем отправленые дынные из формы
 * Если уникальный идентификатор не пустой -
 * редактируем задачу, иначе - создаем
 */
func saveHandler(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    title := r.FormValue("title")
    description := r.FormValue("description")
    var task *models.Task

    if id != "" {
        task = tasks[id]
        task.Title = title
        task.Description = description
    } else {
        id = GenerateId()
        task := models.NewTask(id, title, description)
        tasks[task.Id] = task
    }

    http.Redirect(w, r, "/", 302)
}

/**
 * Метод удаления задачи из ToDoList
 *
 */
func deleteHandler(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    if id == "" {
        http.NotFound(w, r)
    }

    delete(tasks, id)
    http.Redirect(w, r, "/", 302)
}

/**
 * Основной метод ToDoList
 * Метод инициализации программы
 *
 */
func main() {

	/**
	 * Инициализация пустой карты с задачами
	 *
	 */
    tasks = make(map[string]*models.Task, 0)

    /**
	 * Доступ к ресурсам программы
	 *
	 */
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

    /**
	 * Роутинг URL
	 *
	 */
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/create", createHandler)
    http.HandleFunc("/update", updateHandler)
    http.HandleFunc("/delete", deleteHandler)
    http.HandleFunc("/save", saveHandler)

    /**
	 * Старт сервера на 8181-порту
	 *
	 */
    http.ListenAndServe(":8181", nil)
}