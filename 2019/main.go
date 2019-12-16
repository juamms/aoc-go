package main

import (
	"fmt"
	"time"

	aoc "github.com/juamms/aoc-go/2019/days"
	"github.com/juamms/aoc-go/2019/utils"
)

var days = []utils.Day{
	aoc.Day1{}, aoc.Day2{}, aoc.Day3{},
}

var day, part int

func main() {
	utils.ParseFlags(&day, &part)

	if day == 0 {
		runAllDays()
	} else {
		if day < 1 || day > len(days) {
			panic(fmt.Sprintf("%d is not a valid day. (1 <= day <= %d)", day, len(days)))
		}

		if part != 1 && part != 2 {
			part = 0
		}

		if part == 0 {
			runDay(day, 1)
			runDay(day, 2)
		} else {
			runDay(day, part)
		}
	}
}

func runAllDays() {
	start := time.Now()
	for d := 1; d <= len(days); d++ {
		runDay(d, 1)
		runDay(d, 2)
	}

	fmt.Printf("\nTotal elapsed time: %s\n", time.Since(start))
}

func runDay(day, part int) {
	day--

	var result interface{}
	start := time.Now()
	if part == 1 {
		result = days[day].Part1()
	} else {
		result = days[day].Part2()
	}

	day++
	fmt.Printf("[% d|%d]: %v\t(%s)\n", day, part, result, time.Since(start))
}
