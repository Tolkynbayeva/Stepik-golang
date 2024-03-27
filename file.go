package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var arr []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		arr = append(arr, num)
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	sumStr := strconv.Itoa(sum)

	os.Stdout.WriteString(sumStr + "\n")
}
