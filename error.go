package main

func divide(a int, b int) (int, error) {
	return a / b, nil
}

// func main() {
// 	var a, b int
// 	fmt.Scan(&a, &b)
// 	result, err := divide(a, b)
// 	if err != nil {
// 		fmt.Println("ошибка")
// 	} else {
// 		fmt.Println(result)
// 	}
// }
