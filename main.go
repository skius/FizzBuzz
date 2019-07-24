package main

import (
	"fmt"
	"time"
)

func main() {
	const max = 1000

	m := funcTimer(masterString, max)
	d := funcTimer(dynamic, max)
	h := funcTimer(hardcoded, max)

	fmt.Println()
	fmt.Println(m)
	fmt.Println(d)
	fmt.Println(h)
}

func funcTimer(f func(int), max int) int64 {
	start := time.Now().UnixNano()
	f(max)
	end := time.Now().UnixNano()
	return totalTime(start, end)
}

func totalTime(start int64, end int64) int64 {
	return end - start
}
