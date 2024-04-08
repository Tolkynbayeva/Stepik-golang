package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type rat struct {
	Rating []int
}

type st struct {
	Students []rat
}

func main() {
	jsonData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Ошибка при чтении JSON из стандартного ввода:", err)
		return
	}
	var res st
	if err := json.Unmarshal(jsonData, &res); err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}
	count := len(res.Students)
	arr := []int{}

	for _, student := range res.Students {
		arr = append(arr, student.Rating...)
	}

	average := float64(len(arr)) / float64(count)
	result := map[string]float64{"Average": average}
	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при кодировании в JSON:", err)
		return
	}

	fmt.Println(string(jsonResult))
}
