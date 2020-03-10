package main

import (
    "./models"
    "crypto/rand"
    "fmt"
    "html/template"
    "net/http"
)

var tasks map[string]*models.Task

func GenerateId() string {
    b := make([]byte, 16)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("views/index.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    fmt.Println(tasks)
    t.ExecuteTemplate(w, "index", tasks)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("views/form.html", "templates/header.html", "templates/footer.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "form", nil)
}

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

func main() {
    tasks = make(map[string]*models.Task, 0)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/create", createHandler)
    http.HandleFunc("/update", updateHandler)
    http.HandleFunc("/save", saveHandler)
    http.ListenAndServe(":8181", nil)
}