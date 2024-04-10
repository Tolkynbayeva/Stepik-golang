package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func convertUnit(str string) string {
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	formattedTimeStr := parsedTime.Format("Mon Jan _2 15:04:05 -0700 2006")
	return formattedTimeStr
}

func changeDate(str string) string {
	parseTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		panic(err)
	}
	if parseTime.Hour() > 13 {
		parseTime = parseTime.AddDate(0, 0, 1)
	}
	formattedtimeStr := parseTime.Format("2006-01-02 15:04:05")
	return formattedtimeStr
}

func subDuration() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	arr := strings.Split(str, ",")
	str1 := arr[0]
	str2 := arr[1]

	parseTime1, _ := time.Parse("02.01.2006 15:04:05", str1)
	parseTime2, _ := time.Parse("02.01.2006 15:04:05", str2)
	if parseTime1.After(parseTime2) {
		dur := parseTime1.Sub(parseTime2)
		fmt.Println(dur)
	} else {
		dur := parseTime2.Sub(parseTime1)
		fmt.Println(dur)
	}
}

func unixDate(str string) string {
	const now = 1589570165
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	changeInput := strings.ReplaceAll((strings.ReplaceAll(input, " мин. ", "m")), " сек.", "s")
	dur, err := time.ParseDuration(changeInput)
	if err != nil {
		panic(err)
	}
	currentTime := time.Unix(now, 0).UTC()
	res := currentTime.Add(dur)
	currentFormat := res.Format("Mon Jan 2 15:04:05 UTC 2006")
	return currentFormat
}
