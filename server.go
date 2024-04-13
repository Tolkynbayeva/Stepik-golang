package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Напишите веб сервер, который по пути /get отдает текст "Hello, web!".
func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

// Напишите веб-сервер который по пути /api/user приветствует пользователя:
// Принимает и парсит параметр name и делает ответ "Hello,<name>!"
// Пример: /api/user?name=Golang
func handler2(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	message := fmt.Sprintf("Hello,%s!", name)
	w.Write([]byte(message))
}

// Напиши веб сервер (порт :3333) - счетчик который будет обрабатывать GET (/count) и POST (/count) запросы:
// GET:  возвращает счетчик
// POST: увеличивает ваш счетчик на значение  (с ключом "count") которое вы получаете из формы, но если пришло НЕ число то нужно ответить клиенту: "это не число" со статусом http.StatusBadRequest (400).
var count int
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		countStr := r.FormValue("count")
		newCount, err := strconv.Atoi(countStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}

		count += newCount
	}
	fmt.Fprintf(w, "%d", count)
}

func main() {
	http.HandleFunc("/count", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
