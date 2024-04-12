package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// Подключитесь к адресу 127.0.0.1:8081 по протоколу TCP, считайте от сервера 3 сообщения, и выведите их в верхнем регистре. Рекомендуем использовать буфер в 1024 байта.
func connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	message := make([]byte, 1024)
	for i := 0; i < 3; i++ {
		n, err := conn.Read(message)
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.ToUpper(string(message[:n])))
	}
}

// Сделайте HTTP запрос на сервер по пути http://127.0.0.1:5555/get и напечатайте ответ сервера (только тело).
func request() {
	resp, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}

// Считайте с консоли две переменные, сначала name, затем age. Сделайте HTTP запрос на http://127.0.0.1:8080/hello передав name и age в query параметрах, затем напечатайте ответ сервера (только тело).
func query() {
	var name, age string
	fmt.Scan(&name, &age)

	baseURL := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	fullURL := baseURL + "?" + params.Encode()
	resp, err := http.Get(fullURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
