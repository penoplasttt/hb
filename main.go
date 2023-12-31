package main

import "net/http"

func main() {
	// Регистрируем обработчик для статических файлов в директории "static"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Регистрируем обработчик для корневого пути "/"
	http.HandleFunc("/", handler)

	// Запускаем веб-сервер на порту 8080
	http.ListenAndServe(":8080", nil)
}

// handler функция, которая обрабатывает запрос и возвращает HTML-страницу
func handler(w http.ResponseWriter, r *http.Request) {
	// Открываем и читаем HTML-файл
	http.ServeFile(w, r, "static/index.html")
}
