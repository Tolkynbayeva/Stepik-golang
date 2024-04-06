package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	file, _ := os.Open(path)
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil
	}

	if len(data) <= 1 {
		return nil
	}
	fmt.Println("C", data[4][2])
	return nil
}

func main() {
	const root = "."
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v", err)
	}
}
