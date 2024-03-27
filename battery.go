package main

import "fmt"

type Battery struct {
	Capacity string
}

func (b Battery) String() string {
	return fmt.Sprintf("[%v]", b.Capacity)
}

func main() {
	var input string
	fmt.Scan(&input)
	r := []rune(input)
	for i := 0; i < len(r)-1; i++ {
		for j := 0; j < len(r)-i-1; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
	result := ""
	for _, v := range r {
		if v == '0' {
			v = ' '
		} else {
			v = 'X'
		}
		result += string(v)
	}
	batteryForTest := Battery{
		Capacity: result,
	}
	fmt.Println(batteryForTest)
}
