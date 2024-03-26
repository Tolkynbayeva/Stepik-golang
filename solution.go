package main

import (
	"fmt"
	"math"
	"strings"
)

// На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
func Hypot(a float64, b float64) float64 {
	c := (a * a) + (b * b)
	return math.Sqrt(c)
}

// Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
func Star(s string) string {
	n := ""
	for _, v := range s {
		n = n + "*" + string(v)
	}
	n = strings.TrimLeft(n, "*")
	return n
}

// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
func MaxNumber(s string) string {
	max := '0'
	for _, v := range s {
		if v > max {
			max = v
		}
	}

	return string(max)
}

// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число. 
func ChangeNumber(num int) string {
	n := 0
	result := ""
	for num > 0 {
		n = num % 10
		num /= 10
		result = fmt.Sprint(int(n*n)) + result
	}

	return result
}


// Требуется вычислить период колебаний (t) 
var k, p, v float64 = 1296, 6, 6

func M() float64 {
	m := p * v
	return m
}

func W() float64 {
	w := math.Sqrt(k / M())
	return w
}

func T() float64 {
	t := 6 / W()
	return t
}