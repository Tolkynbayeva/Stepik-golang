package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(n uint) uint {
		if n == 0 {
			return uint(100)
		}
		str := ""
		for n > 0 {
			num := n % 10
			if num%2 == 0 && num != 0 {
				str = strconv.Itoa(int(num)) + str
			}
			n /= 10
		}
		if str == "" {
			return 100
		}
		res, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			panic(err)
		}
		return uint(res)
	}
	fmt.Println(fn(71))
}
