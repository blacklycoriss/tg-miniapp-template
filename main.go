package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Проверяем, существует ли файл index.html
	if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		log.Fatal("Файл index.html не найден в текущей директории")
	}

	// Регистрируем обработчик для корневого пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Только точный корень, без подпути
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Отдаём файл index.html
		http.ServeFile(w, r, "index.html")
	})

	// Можно также просто использовать FileServer (альтернативный вариант ниже в комментариях)

	fmt.Println("Сервер запущен → http://localhost:8443")
	fmt.Println("Отдаётся файл: index.html")

	// Запускаем сервер на порту 8080
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
