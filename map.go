package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	ns := strings.ReplaceAll((strings.ReplaceAll(s, ",", ".")), " ", "")
	arr := strings.Split(ns, ";")
	num1, err := strconv.ParseFloat(arr[0], 64)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		panic(err)
	}
	div := num1 / num2
	fmt.Printf("%.4f", div)
}

// 1 149,6088607594936;1 179,0666666666666