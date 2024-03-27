package main

import (
	"fmt"
)

// только для проверки
func readTask() (interface{}, interface{}, interface{}) {
	return 6.0, 4.7, "+"
}

func main() {
	value1, value2, operation := readTask()
	if _, ok := value1.(float64); !ok {
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}
	if _, ok := value2.(float64); !ok {
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}

	op, ok := operation.(string)
	if !ok {
		fmt.Println("неизвестная операция")
		return
	}
	var result float64
	switch op {
	case "+":
		result = value1.(float64) + value2.(float64)
	case "-":
		result = value1.(float64) - value2.(float64)
	case "*":
		result = value1.(float64) * value2.(float64)
	case "/":
		if value2.(float64) == 0 {
			fmt.Printf("value=%v: %T", value1, value1)
			return
		}
		result = value1.(float64) / value2.(float64)
	default:
		fmt.Println("неизвестная операция")
		return
	}
	fmt.Printf("%.4f", result)
}
