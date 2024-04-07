package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	nums := []string{}
	for {
		data, _ := reader.ReadString(';')
		if data == "" {
			break
		}
		nums = append(nums, data)
	}
	res := 1
	for i, v := range nums {
		if v == "0;" {
			res = i + res
		}
	}
	fmt.Println(res)
}
