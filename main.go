package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

// Данные о состоянии Кирилла
type PageData struct {
	Name   string
	Weight int
	Size   int
}

var (
	weight = 100 // Начальный вес
	size   = 150 // Начальный размер картинки в пикселях
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/feed", feedHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").Parse(htmlTemplate))
	mu.Lock()
	data := PageData{Name: "Кирилл Свинота", Weight: weight, Size: size}
	mu.Unlock()
	tmpl.Execute(w, data)
}

func feedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		mu.Lock()
		weight += 10
		size += 15 // Увеличиваем размер, чтобы он становился толще
		mu.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// HTML-шаблон прямо в коде для удобства
const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Кормушка</title>
    <style>
        body { text-align: center; font-family: sans-serif; background-color: #f0f0f0; }
        .kirill { 
            transition: all 0.3s ease; 
            display: inline-block;
            background-color: pink;
            border-radius: 50%;
            border: 5px solid #ff80ab;
            margin: 20px;
        }
        button { padding: 10px 20px; font-size: 20px; cursor: pointer; background: #4CAF50; color: white; border: none; border-radius: 5px; }
        button:hover { background: #45a049; }
    </style>
</head>
<body>
    <h1>{{.Name}}</h1>
    <p>Вес: {{.Weight}} кг</p>
    
    <div class="kirill" style="width: {{.Size}}px; height: {{.Size}}px; line-height: {{.Size}}px;">
        🐷
    </div>

    <form action="/feed" method="post">
        <button type="submit">Покормить Кирилла 🍎</button>
    </form>
</body>
</html>
`
