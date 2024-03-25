package main

import (
	"strings"
)

func correct(s string) string {
	r := []rune(s)
	c := ""
	if (r[0] >= 'А' && r[0] <= 'Я') && r[len(r)-1] == '.' {
		c = "Right"
	} else {
		c = "Wrong"
	}
	return c
}

func palindrom(s string) string {
	r := ""
	p := "Палиндром"
	for _, v := range s {
		r = string(v) + r
	}
	if r != s {
		p = "Нет"
	}
	return p
}

func index(x string, s string) int {
	count := strings.Index(x, s)
	return count
}

func removeLetter(s string) string {
	n := ""
	for i, v := range s {
		if i%2 != 0 {
			n = n + string(v)
		}
	}
	return n
}

func removeRepeat(s string) string {
	for _, v := range s {
		if strings.Count(s, string(v)) > 1 {
			s = strings.ReplaceAll(s, string(v), "")
		}
	}
	return s
}

func checkPassword(s string) string {
	res := "Wrong password"
	if len(s) >= 5 {
		for _, v := range s {
			if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
				res = "Ok"
			} else {
				res = "Wrong password"
				break
			}
		}
	}
	return res
}